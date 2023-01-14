package rock_paper_scissors

import "testing"

func TestRPSFullPlaysScore(t *testing.T) {
	type test struct {
		opp_plays  []Hand
		your_plays []Hand
		want       int
	}

	example_opp_plays := []Hand{Rock, Paper, Scissors}
	example_your_plays := []Hand{Paper, Rock, Scissors}

	tests := []test{
		{opp_plays: example_opp_plays, your_plays: example_your_plays, want: 15},
	}

	for _, tc := range tests {
		got := RPSFullPlaysScore(tc.opp_plays, tc.your_plays)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v\n", tc.want, got)
		}
	}

}

func TestRPSPlayShapeScore(t *testing.T) {
	type test struct {
		opp_plays []Hand
		outcomes  []Outcome
		want      int
	}

	example_opp_plays := []Hand{Rock, Paper, Scissors}
	example_outcome := []Outcome{Draw, Lose, Win}

	tests := []test{
		{opp_plays: example_opp_plays, outcomes: example_outcome, want: 12},
	}

	for _, tc := range tests {
		got := RPSPlayShapeScore(tc.opp_plays, tc.outcomes)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v\n", tc.want, got)
		}
	}

}
