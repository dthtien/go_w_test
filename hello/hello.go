package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name, lang string) string {
  if name == "" { name = "World" }

  return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
  switch lang {
  case spanish:
    prefix = spanishHelloPrefix
  case french:
    prefix = frenchHelloPrefix
  default:
    prefix = englishHelloPrefix
  }
  return
}

func main() {
  fmt.Println(Hello("Chris", "English"))
}
