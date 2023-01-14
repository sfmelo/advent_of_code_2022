package calorie_counting

import (
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
