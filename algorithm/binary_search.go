package main

import (
	"fmt"
)

func main() {
	sortNum := []int{1, 2, 5, 6, 7, 9}
	selectNnum := Search(5, sortNum)
	fmt.Println(selectNnum)
}

func Search(targetNum int, sortArray []int) int {
	left := 0
	rigth := len(sortArray) - 1

	for left <= rigth {
		mid := (left + rigth) / 2
		if targetNum == sortArray[mid] {
			return sortArray[mid]
		}
		if targetNum > sortArray[mid] {
			left = mid + 1
			continue
		}
		if targetNum <= sortArray[mid] {
			rigth = mid - 1
			continue
		}
	}

	return -1
}
