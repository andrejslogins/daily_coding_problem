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
	k := readK()

	val := false
	for i, s := range arr {
		for _, t := range arr[i:] {
			if (s + t) == k {
				val = true
			}
		}
	}
	printSlice(arr)
	fmt.Printf("k: %v\n", k)
	fmt.Printf("Result: %v\n", val)
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

func readK() int {
	fmt.Println("Input k:")
	var s int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		val, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Must be an integer")
			continue
		}
		s = val
		break
	}
	return s
}

// Good morning! Here's your coding interview problem for today.
// This problem was recently asked by Google.
// Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
// Bonus: Can you do this in one pass?
