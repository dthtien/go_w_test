package main

type Dictionary map[string]string

const (
  WordNotFound   = DictionaryErr("could not find the word you were looking for")
  WordExisted = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}
func (d Dictionary) Search(word string) (string, error){
  found, ok := d[word]

  if !ok {
    return "", WordNotFound
  }

  return found, nil
}

func (d Dictionary) Update(word, definition string) error {
  _, err := d.Search(word)

  switch err {
  case WordNotFound:
    return WordNotFound
  case nil:
    d[word] = definition
  default:
    return err
  }

  return nil
}

func (d Dictionary) Delete(word string) error {
  _, err := d.Search(word)

  switch err {
  case WordNotFound:
    return WordNotFound
  case nil:
    delete(d, word)
  default:
    return err
  }

  return nil
}

func (d Dictionary) Add(word, definition string) error {
  _, err := d.Search(word)

  switch err {
  case WordNotFound:
    d[word] = definition
  case nil:
    return WordExisted
  default:
    return err
  }

  return nil
}

