package main

import (
	"advent-of-code-2023/lib"
	"strconv"
	"strings"
)

const SmallTestString string = ``

const TestString string = `Time:      7  15   30
Distance:  9  40  200`

const DataFile string = "data.txt"

/*
	Part 1 Notes

	For each game we have to find out how many different ways there are to beat the best score
	Then we multiply them together

	To do this let's start by defining the games with their time and their current records

*/

type Game struct {
	Time             int
	Record           int
	PossibleOutcomes []Outcome
}

type Outcome struct {
	HeldFor      int
	TravelledFor int
	Distance     int
}

func makeGames(input string) (games []Game) {
	times := lib.IntsFromString(getTimesNumbers(input))
	records := lib.IntsFromString(getRecordsNumbers(input))

	for i := 0; i < len(times); i += 1 {
		newGame := Game{Time: times[i], Record: records[i]}
		newGame.PossibleOutcomes = determineOutcomes(newGame)
		games = append(games, newGame)
	}

	return
}

func getTimesNumbers(input string) string {
	lines := strings.Split(input, "\n")
	return strings.Split(lines[0], ":")[1]
}

func getRecordsNumbers(input string) string {
	lines := strings.Split(input, "\n")
	return strings.Split(lines[1], ":")[1]
}

func determineOutcomes(game Game) (outcomes []Outcome) {
	for i := 1; i <= game.Time; i++ {
		heldFor := i
		travelledFor := game.Time - i
		distance := heldFor * travelledFor
		newOutcome := Outcome{HeldFor: heldFor, TravelledFor: travelledFor, Distance: distance}
		outcomes = append(outcomes, newOutcome)
	}
	return
}

func numberOfWinningOutcomes(game Game) (number int) {
	for _, outcome := range game.PossibleOutcomes {
		if outcome.Distance > game.Record {
			number += 1
		}
	}
	return
}

func solvePart1(input string) int {
	games := makeGames(input)
	result := lib.Reduce(games, func(sum int, game Game) int {
		return sum * numberOfWinningOutcomes(game)
	}, 1)

	return result
}

/*
	Part 2 Notes

	So input is combining here: 7  15   30 becomes 71530
	First order of business is fixing the input
	Then I should be able to run almost the same code.
	I could make it more performant, but probably don't need to
	I could make it more performant by only calcualating up to the halfway mark

*/

func makeGame(input string) Game {
	timeString := strings.ReplaceAll(getTimesNumbers(input), " ", "")
	recordString := strings.ReplaceAll(getRecordsNumbers(input), " ", "")
	time, _ := strconv.Atoi(timeString)
	record, _ := strconv.Atoi(recordString)

	game := Game{Time: time, Record: record}
	game.PossibleOutcomes = determineOutcomes(game)

	return game
}

func solvePart2(input string) int {
	game := makeGame(input)
	return numberOfWinningOutcomes(game)
}

func main() {
	lib.AssertEqual(288, solvePart1(TestString))
	lib.AssertEqual(71503, solvePart2(TestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
