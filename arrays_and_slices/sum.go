package main

func Sum(numbers []int) int {
  sum := 0
  for _, numb := range numbers { sum += numb }
  return sum
}

func SumAll(numberstosum ...[]int) (sums []int) {
  for _, numbs := range numberstosum {
    sums = append(sums, Sum(numbs))
  }

  return
}

func SumAllTails(numberstosum ...[]int) (sums []int) {
  for _, numbs := range numberstosum {
    if len(numbs) == 0 {
      sums = append(sums, 0)
    } else {
      tail := numbs[1:]
      sums = append(sums, Sum(tail))
    }
  }

  return
}

