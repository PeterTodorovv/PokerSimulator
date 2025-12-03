package utils

import "sort"

func CalculateHand(hand Hand, river []Card) HandResult {
	scoreHand := append(hand.cards, river...)
	sortedHand := sortHand(scoreHand)
	combination := calculateCombinations2(hand, river)

	straightCombination := getStreight(sortedHand)

	if straightCombination.Combination != -1 && combination.Combination < STRAIGHT {
		combination = straightCombination
	}

	flushCombination := getFlush(sortedHand)

	if flushCombination.Combination != -1 && combination.Combination < FLUSH {
		combination = flushCombination

		//TODO Fix royal flush

		straightFlushResult := getStreightFlush(sortedHand)

		if straightFlushResult.Combination != -1 {
			straightFlushResult.Combination = STRAIGHT_FLUSH
			combination = straightFlushResult
		}
	}

	return combination
}

func CompareHands(hands []Hand, river []Card) GameResult {
	gameResult := GameResult{Winner: -1}
	handResults := []HandResult{}
	pointer := 0
	for _, hand := range hands {
		handResults = append(handResults, CalculateHand(hand, river))
	}

	for i := 1; i < len(hands); i++ {
		hand1Result := handResults[pointer]
		hand2Result := handResults[i]

		if hand1Result.Combination > hand2Result.Combination {
			gameResult.Winner = pointer
		} else if hand1Result.Combination < hand2Result.Combination {
			pointer = i
			gameResult.Winner = pointer
		} else {
			activeCards1 := hand1Result.acticveCards
			activeCards2 := hand2Result.acticveCards

			for j := 0; j < 5; j++ {
				if activeCards1[j] > activeCards2[j] {
					gameResult.Winner = pointer
					break
				} else if activeCards1[j] < activeCards2[j] {
					gameResult.Winner = i
					pointer = i
					break
				}
			}
		}
	}

	return gameResult
}

func calculateCombinations2(hand Hand, river []Card) HandResult {
	scoreHand := append(hand.cards, river...)
	sortedHand := sortHand(scoreHand)

	numbers := map[int]int{}
	result := HandResult{Combination: 0}

	for _, card := range sortedHand {
		_, ok := numbers[card.Number]
		if ok == false {
			numbers[card.Number] = 0
		}

		numbers[card.Number]++
	}

	sortedKv := sortMap(numbers)
	scoringNumbers := []int{}

	for _, kv := range sortedKv {

		if kv.Value == 4 {
			result.Combination = FOUR_OF_A_KIND

			for i := 0; i < 4; i++ {
				scoringNumbers = append(scoringNumbers, kv.Key)
			}

			for _, card := range sortedHand {
				if card.Number != kv.Key {
					scoringNumbers = append(scoringNumbers, card.Number)

					if len(scoringNumbers) == 5 {
						break
					}
				}
			}
		}

		if kv.Value == 3 {
			result.Combination = THREE_OF_A_KIND

			for i := 0; i < 3; i++ {
				scoringNumbers = append(scoringNumbers, kv.Key)
			}
		}

		if kv.Value == 2 {

			if result.Combination == 0 {
				result.Combination = PAIR
			} else if result.Combination == 1 {
				result.Combination = TWO_PAIR
			} else if result.Combination == 3 {
				result.Combination = FULL_HOUSE
			}

			for i := 0; i < kv.Value; i++ {
				scoringNumbers = append(scoringNumbers, kv.Key)
			}

		}

		if kv.Value == 1 {
			for i := 0; i < kv.Value; i++ {
				scoringNumbers = append(scoringNumbers, kv.Key)
			}

		}

		if len(scoringNumbers) == 5 {
			break
		}
	}

	result.acticveCards = scoringNumbers
	return result
}

func sortMap(m map[int]int) []kv {

	// Convert map entries to a slice

	var pairs []kv
	for k, v := range m {
		pairs = append(pairs, kv{k, v})
	}

	// Sort by value desc, then key desc
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Value == pairs[j].Value {
			return pairs[i].Key > pairs[j].Key // key descending
		}
		return pairs[i].Value > pairs[j].Value // value descending
	})

	return pairs
}

func sortHand(cards []Card) []Card {
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			if cards[i].Number < cards[j].Number {
				tmp := cards[i]
				cards[i] = cards[j]
				cards[j] = tmp
			}
		}
	}

	return cards
}

func getFlush(scoreHand []Card) HandResult {
	clubs := []int{}
	diamons := []int{}
	hearts := []int{}
	spades := []int{}

	for _, card := range scoreHand {
		switch card.Suit {
		case CLUBS:
			clubs = append(clubs, card.Number)
		case DIAMONDS:
			diamons = append(diamons, card.Number)
		case HEARTS:
			hearts = append(hearts, card.Number)
		case SPADES:
			spades = append(spades, card.Number)
		}
	}

	if len(clubs) == 5 {
		return HandResult{
			Combination:  FLUSH,
			acticveCards: clubs,
		}
	}

	if len(diamons) == 5 {
		return HandResult{
			Combination:  FLUSH,
			acticveCards: diamons,
		}
	}

	if len(hearts) == 5 {
		return HandResult{
			Combination:  FLUSH,
			acticveCards: hearts,
		}
	}

	if len(spades) == 5 {
		return HandResult{
			Combination:  FLUSH,
			acticveCards: spades,
		}
	}

	return HandResult{
		Combination: -1,
	}
}

func getStreight(scoreHand []Card) HandResult {
	hasAce := scoreHand[0].Number == 14
	lastCard := 0
	currentStreak := []Card{}
	highStreak := []Card{}

	for i := 0; i < len(scoreHand); i++ {
		if scoreHand[i].Number == lastCard {
			continue
		}

		if lastCard-scoreHand[i].Number == 1 {
			currentStreak = append(currentStreak, scoreHand[i])

			if len(currentStreak) > len(highStreak) {
				highStreak = currentStreak
			}
		} else {
			currentStreak = []Card{scoreHand[i]}
		}

		lastCard = scoreHand[i].Number
	}

	if hasAce && len(highStreak) == 4 && highStreak[0].Number == 5 {
		highStreak = append(highStreak, scoreHand[0])
	}

	straightNumbers := []int{}

	for _, card := range highStreak {
		straightNumbers = append(straightNumbers, card.Number)
	}
	if len(highStreak) >= 5 {
		return HandResult{
			Combination:  STRAIGHT,
			acticveCards: straightNumbers,
		}
	}

	return HandResult{
		Combination: -1,
	}
}

func getStreightFlush(scoreHand []Card) HandResult {
	hasAce := scoreHand[0].Number == 14
	lastCard := 0
	currentSuit := -1
	currentStreak := []Card{}
	highStreak := []Card{}

	for i := 0; i < len(scoreHand); i++ {
		if scoreHand[i].Number == lastCard {
			continue
		}

		if lastCard-scoreHand[i].Number == 1 && currentSuit == scoreHand[i].Suit {
			currentStreak = append(currentStreak, scoreHand[i])

			if len(currentStreak) > len(highStreak) {
				highStreak = currentStreak
			}
		} else {
			currentStreak = []Card{scoreHand[i]}
		}

		lastCard = scoreHand[i].Number
		currentSuit = scoreHand[i].Suit

	}

	if hasAce && len(highStreak) == 4 && highStreak[0].Number == 5 && highStreak[0].Suit == scoreHand[0].Suit {
		highStreak = append(highStreak, scoreHand[0])
	}

	straightNumbers := []int{}

	for _, card := range highStreak {
		straightNumbers = append(straightNumbers, card.Number)
	}
	if len(highStreak) >= 5 {
		return HandResult{
			Combination:  STRAIGHT,
			acticveCards: straightNumbers,
		}
	}

	return HandResult{
		Combination: -1,
	}
}
func getCombinationLengthMap() CombinationLength {
	return CombinationLength{
		1: HIGH_CARD,
		2: PAIR,
		3: THREE_OF_A_KIND,
		4: FOUR_OF_A_KIND,
	}
}

//add ordering for the hand, so AQ will be the same as QA

type HandResult struct {
	Combination         int
	highCardCombination int
	acticveCards        []int
}

type GameResult struct {
	Winner int
}
type CombinationLength map[int]int

type kv struct {
	Key   int
	Value int
}
