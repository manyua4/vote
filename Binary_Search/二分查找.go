package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1

		} else {
			high = mid - 1
		}
	}
	return -1
}
func main() {
	arr := []int{1, 3, 4, 6, 8, 9, 10, 11}
	target := 10
	index := BinarySearch(arr, target)
	if index != -1 {
		fmt.Printf("目标元素%d的位置为%d", target, index)
	} else {
		fmt.Println("查无此数")
	}
}
