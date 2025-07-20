package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.

func FirstTurn(card1, card2, dealerCard string) string {
	player1 := ParseCard(card1)
	player2 := ParseCard(card2)
	dealer := ParseCard(dealerCard)

	handValue := player1 + player2

	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	if handValue == 21 {
		if dealer >= 10 {
			return "S"
		} else {
			return "W" // Win immediately
		}
	}

	if handValue >= 17 && handValue <= 20 {
		return "S"
	}

	if handValue >= 12 && handValue <= 16 {
		if dealer >= 7 {
			return "H"
		} else {
			return "S"
		}
	}

	return "H"
}
