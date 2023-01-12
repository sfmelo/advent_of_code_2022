package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func CalorieCounting(input string, top int) int {
	calories := strings.Split(input, "\n")

	// convert string of calories into slice of each elf calorie sum
	var elvesCalories []int
	calorieSum := 0
	for _, c := range calories {
		calories, err := strconv.Atoi(c)

		if err != nil {
			elvesCalories = append(elvesCalories, calorieSum)
			calorieSum = 0
		}

		calorieSum += calories
	}
	elvesCalories = append(elvesCalories, calorieSum)

	// sort elf calories in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(elvesCalories)))

	// sum of top elf calories
	sum := 0
	for i := 0; i < top; i++ {
		sum += elvesCalories[i]
	}

	return sum
}

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("reading file %v: %v", "input.txt", err)
		return
	}

	fmt.Printf("Solution of top 1 is: %v\n", CalorieCounting(string(buf), 1))
	fmt.Printf("Solution of top 3 is: %v\n", CalorieCounting(string(buf), 3))

}
