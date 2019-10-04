// As we iterate through the list, the array becomes sorted 1 element at a time
func insertionSort(array []int) []int{
  for i := 0; i < len(array); i++{
    j := i // start at that element, compare prior
    for (j > 0) && (array[j-1]>array[j]){ // While not end of the array, and element prior is bigger
      temp := array[j] // Swap element with the one to its left
      array[j] = array[j-1]
      array[j-1] = temp
      j -= 1
    }
  }
  return array
}
