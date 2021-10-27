package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"strings"
)

const (
  titleSeparator       = "Title: "
  descriptionSeparator = "Description: "
  tagsSeparator         = "Tags: "
)


func newPost(postFile fs.File) (Post, error) {
  scanner := bufio.NewScanner(postFile)

  readMetaLine := func(tagName string) string {
    scanner.Scan()
    return strings.TrimPrefix(scanner.Text(), tagName)
  }

  title := readMetaLine(titleSeparator)
  description := readMetaLine(descriptionSeparator)
  tags := strings.Split(readMetaLine(tagsSeparator), ", ")

  post := Post{
    Title: title,
    Description: description,
    Tags: tags,
    Body: readBody(scanner),
  }
  return post, nil
}

func readBody(scanner *bufio.Scanner) string {
  scanner.Scan()
  buf := bytes.Buffer{}
  for scanner.Scan() {
    fmt.Fprintln(&buf, scanner.Text())
  }

  return strings.TrimSuffix(buf.String(), "\n")
}
