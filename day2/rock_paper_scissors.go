package rock_paper_scissors

type Hand string

const (
	Rock     Hand = "Rock"
	Paper    Hand = "Hand"
	Scissors Hand = "Scissors"
)

type Outcome int

const (
	Win  Outcome = 6
	Lose Outcome = 0
	Draw Outcome = 3
)

func RPSFullPlaysScore(opp_plays []Hand, your_plays []Hand) int {
	score := 0

	// Assume len(opp_plays) == len(your_plays)
	for i := 0; i < len(opp_plays); i++ {
		score += shapeScore(your_plays[i]) + int(outcomeScore(opp_plays[i], your_plays[i]))
	}

	return score
}

func RPSPlayShapeScore(opp_plays []Hand, outcomes []Outcome) int {
	score := 0

	// Assume len(opp_plays) == len(outcomes)
	for i := 0; i < len(opp_plays); i++ {
		your_hand := neededHand(opp_plays[i], outcomes[i])
		score += shapeScore(your_hand) + int(outcomeScore(opp_plays[i], your_hand))
	}

	return score
}

func shapeScore(h Hand) int {
	switch h {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		return 0
	}
}

func outcomeScore(opponent Hand, you Hand) Outcome {
	if opponent == you {
		return Draw
	}

	if you == Rock && opponent == Scissors ||
		you == Paper && opponent == Rock ||
		you == Scissors && opponent == Paper {
		return Win
	}

	return Lose
}

func neededHand(opponent Hand, outcome Outcome) Hand {
	switch outcome {
	case Win:
		switch opponent {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	case Lose:
		switch opponent {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}
	}

	// in case of draw
	return opponent
}
