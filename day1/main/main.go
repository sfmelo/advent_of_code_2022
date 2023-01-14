package main

import (
	"calorie_counting"
	"fmt"
	"io/ioutil"
)

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("reading file %v: %v", "input.txt", err)
		return
	}

	fmt.Printf("Solution of top 1 is: %v\n", calorie_counting.CalorieCounting(string(buf), 1))
	fmt.Printf("Solution of top 3 is: %v\n", calorie_counting.CalorieCounting(string(buf), 3))

}
