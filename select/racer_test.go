package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
  return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    time.Sleep(delay)
    w.WriteHeader(http.StatusOK)
  }))

}

func TestRacer(t *testing.T) {
  t.Run("returns correct url", func(t *testing.T) {
    slowServer := makeDelayedServer(20 * time.Millisecond)
    fastServer := makeDelayedServer(0 * time.Millisecond)

    defer slowServer.Close()
    defer fastServer.Close()

    fastURL := fastServer.URL
    slowURL := slowServer.URL

    want := fastURL
    got, err := Racer(slowURL, fastURL)

    if got != want {
      t.Errorf("got %q, want %q", got, want)
    }

    if err != nil {
      t.Fatalf("did not expect an error but got one %v", err)
    }
  })

  t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
    server := makeDelayedServer(25 * time.Millisecond)

    defer server.Close()

    _, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

    if err == nil {
      t.Error("expected an error but didn't get one")
    }
  })
}
