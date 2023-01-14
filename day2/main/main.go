package main

import (
	"fmt"
	"io/ioutil"
	rps "rock_paper_scissors"

	"strings"
)

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("reading file %v: %v", "input.txt", err)
		return
	}
	input := string(buf)

	var opp_plays []rps.Hand
	var your_plays []rps.Hand
	var outcomes []rps.Outcome

	for _, line := range strings.Split(input, "\n") {
		plays := strings.Split(line, " ")
		if len(plays) != 2 {
			continue
		}
		// get opponent hands
		switch plays[0] {
		case "A":
			opp_plays = append(opp_plays, rps.Rock)
		case "B":
			opp_plays = append(opp_plays, rps.Paper)
		case "C":
			opp_plays = append(opp_plays, rps.Scissors)
		}
		// get your hands
		switch plays[1] {
		case "X":
			your_plays = append(your_plays, rps.Rock)
			outcomes = append(outcomes, rps.Lose)
		case "Y":
			your_plays = append(your_plays, rps.Paper)
			outcomes = append(outcomes, rps.Draw)
		case "Z":
			your_plays = append(your_plays, rps.Scissors)
			outcomes = append(outcomes, rps.Win)
		}
	}

	fmt.Printf("Solution when second column are your plays: %v\n", rps.RPSFullPlaysScore(opp_plays, your_plays))
	fmt.Printf("Solution when second column are outcomes: %v\n", rps.RPSPlayShapeScore(opp_plays, outcomes))

}
