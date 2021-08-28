package iteration

const repeatedCount = 5

func Repeat(character string) (text string) {
  for i := 0; i < repeatedCount; i ++ {
    text += character
  }

  return
}

