package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printSlice(s []int) {
	fmt.Printf("List: %v\n", s)
}
func main() {
	arr := readIntegerList()
	res := calculateNewArray(arr)

	printSlice(arr)
	printSlice(res)
}

func readIntegerList() []int {
	fmt.Println("Input the list:")
	var s []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		val, err := strconv.Atoi(text)
		if err != nil {
			if text != "" {
				fmt.Println("Must be an integer")
				continue
			} else {
				break
			}
		}
		s = append(s, val)
		printSlice(s)
	}

	return s
}

func calculateNewArray(arr []int) []int {
	var newArr []int
	for i := len(arr) - 1; i >= 0; i-- {
		var initial = true
		var elem int
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			}
			if initial {
				elem = arr[j]
				initial = false
			} else {
				elem = elem * arr[j]
			}
		}
		newArr = append(newArr, elem)
	}
	return newArr
}

// Good morning! Here's your coding interview problem for today.
// This problem was asked by Uber.
// Given an array of integers, return a new array such that each element at index i of the new array is the product
// of all the numbers in the original array except the one at i.
// For example, if our input was [1, 2, 3, 4, 5], the expected output would be
// [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].
