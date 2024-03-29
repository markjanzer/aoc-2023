package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"strconv"
	"strings"
)

const SmallTest1 string = `???.### 1,1,3`           // => 1
const SmallTest2 string = `.??..??...?##. 1,1,3`    // => 4 | 16384 (part 2)
const SmallTest3 string = `?#?#?#?#?#?#?#? 1,3,1,6` // => 1
const SmallTest4 string = `????.#...#... 4,1,1`     // => 16 (part 2)

const TestString string = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

const DataFile string = "data.txt"

/*
	Part 1 Notes

	Okay first we get the row and the groups (of damaged springs), and we pass that to
	Then we iterate through the string with a solve method
	First this method checks if the given string has more damaged springs than provided in
		the groups, and if it does then it returns 0
	This method iterates over the string. When it comes across a ?
		It replaces the character with a . and with a #, and it returns the sum of solve being called
		with those two strings
	If it is not a ? then continue.
	It returns 1

*/

// Takes a line and returns a map of the springs, and a log of the contiguous damaged springs
func getMapAndLog(input string) ([]byte, []int) {
	parts := strings.Split(input, " ")
	springsMap := []byte(parts[0])
	damagedSpringsLog := lib.Map(strings.Split(parts[1], ","), func(item string) int {
		num, _ := strconv.Atoi(item)
		return num
	})

	return springsMap, damagedSpringsLog
}

var cache map[string]int = make(map[string]int)

func serialize(springsMap []byte, damagedSpringsCount int, groups []int) string {
	return string(springsMap) + ":" + fmt.Sprint(damagedSpringsCount) + ":" + strings.Join(lib.Map(groups, func(item int) string {
		return strconv.Itoa(item)
	}), ",")
}

func solveWithCache(springsMap []byte, damagedSpringsCount int, groups []int) int {
	serialized := serialize(springsMap, damagedSpringsCount, groups)
	if _, ok := cache[serialized]; !ok {
		cache[serialized] = solve(springsMap, damagedSpringsCount, groups)
	}
	return cache[serialized]
}

func solve(springsMap []byte, damagedSpringsCount int, groups []int) int {
	if len(groups) == 0 {
		if hasDamagedSpring(springsMap) {
			return 0
		} else {
			return 1
		}
	}

	if len(springsMap) == 0 {
		return 0
	}

	contiguousDamagedSprings := groups[0]

	for i, b := range springsMap {
		if string(b) == "." && damagedSpringsCount > 0 {
			if damagedSpringsCount == contiguousDamagedSprings {
				return solveWithCache(springsMap[(i+1):], 0, groups[1:])
			} else {
				return 0
			}
		} else if string(b) == "#" {
			damagedSpringsCount++
		} else if string(b) == "?" {
			mapWithUndamagedSpring := append([]byte{lib.CharToByte(".")}, springsMap[i+1:]...)
			mapWithDamagedSpring := append([]byte{lib.CharToByte("#")}, springsMap[i+1:]...)
			return solveWithCache(mapWithUndamagedSpring, damagedSpringsCount, groups) + solveWithCache(mapWithDamagedSpring, damagedSpringsCount, groups)
		}
	}

	if damagedSpringsCount == contiguousDamagedSprings {
		return solveWithCache([]byte{}, 0, groups[1:])
	} else {
		return 0
	}
}

func hasDamagedSpring(springsMap []byte) bool {
	return lib.Any(springsMap, func(spring byte) bool {
		return string(spring) == "#"
	})
}

func solvePart1(input string) (sum int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		row, groups := getMapAndLog(line)
		sum += solveWithCache(row, 0, groups)
	}
	return
}

/*
	Part 2 Notes

	Make part 1 more efficient
	Okay there are two steps to this. Get a new getMapAndLog method that multiplies the results by 5

	How do I make part 1 more efficient?
	First off let's get some better naming to make this simpler. I like calling one part the map, and the other the
	log. Like logOfContiguousDamagedSprings.

	Okay instead of creating all of the possible maps and then comparing them to the log,
	let's iterate through the maps, but shorten the map and log when possible. Return 0 when it breaks or doesn't match
	and then return 1 when they are both empty.
*/

func multiplyRowsAndGroups(multiplier int, springsMap []byte, springsLog []int) (newSpringsMap []byte, newSpringsLog []int) {
	for i := 0; i < FoldMultiplier; i++ {
		newSpringsMap = append(newSpringsMap, springsMap...)
		newSpringsLog = append(newSpringsLog, springsLog...)

		if i != multiplier-1 {
			newSpringsMap = append(newSpringsMap, lib.CharToByte("?"))
		}
	}
	return
}

const FoldMultiplier int = 5

func solveLine2(line string) int {
	springsMap, springsLog := getMapAndLog(line)
	springsMap, springsLog = multiplyRowsAndGroups(FoldMultiplier, springsMap, springsLog)
	return solveWithCache(springsMap, 0, springsLog)
}

func solvePart2(input string) (sum int) {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		fmt.Println("Solving line", i+1, "of", len(lines))
		sum += solveLine2(line)
	}
	return
}

func main() {
	lib.AssertEqual(21, solvePart1(TestString))

	// This was breaking because my serialization wasn't separating numbers with commas
	// so 1,1 was being treated as 11 and vice versa
	lib.AssertEqual(1, solvePart1("#?#??..?.#?????#???# 1,1,1,1,11"))

	lib.AssertEqual(1, solveLine2(SmallTest1))
	lib.AssertEqual(16384, solveLine2(SmallTest2))
	lib.AssertEqual(1, solveLine2(SmallTest3))
	lib.AssertEqual(16, solveLine2(SmallTest4))
	lib.AssertEqual(525152, solvePart2(TestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
