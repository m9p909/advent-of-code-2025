package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	data := string(content)

	fmt.Printf("%d", solveProblem2(data))

}

func solveProblem(data string) int {
	// count number of times the dial is rotated to 0
	// start at 50
	// start at 50
	// maintain sum mod 100, count number of times sum gets to 0

	// parse data
	d := strings.Split(data, "\n")
	sum := 50
	total := 0

	for _, v := range d {
		if len(v) <= 0 {
			continue
		}
		dir := v[0]
		count, _ := strconv.Atoi(v[1:])
		if dir == 'L' {
			sum = (sum - count) % 100
		}
		if dir == 'R' {
			sum = (sum + count) % 100
		}
		if sum == 0 {
			total++
		}
	}
	return total

}

func solveProblem2(data string) int {
	// count number of times the dial passes 0
	// start at 50
	// start at 50
	// When it passes 0, then on the left side sum - count is less than 0
	// when it passes 0, then on the right sum + count is greater than 100
	// possible for one rotation to pass 0 multiple times!
	// count number divided by 100 + 1 to be the number of times 0 is passed
	// positive rot c = rot/100
	// negative rot c = -rot/100 +1
	// prev !=0 is not going to work for case hwere
	// 0 then L10000
	// if prev = 0 then total - 1

	// parse data
	d := strings.Split(data, "\n")
	sum := 50
	total := 0
	prev := 50

	for _, v := range d {
		if len(v) <= 0 {
			continue
		}
		dir := v[0]
		count, _ := strconv.Atoi(v[1:])
		if dir == 'L' {
			sum = (sum - count)
		}
		if dir == 'R' {
			sum = (sum + count)
		}
		if sum >= 100 {
			total += sum / 100
		}
		if sum <= 0 {
			total += -sum/100 + 1
			if prev == 0 {
				total -= 1

			}
		}
		sum = mod(sum, 100)
		prev = sum
		fmt.Printf("input: %s, sum: %d, total %d \n", v, sum, total)
	}
	return total

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func mod(a, b int) int {
	return (a%b + b) % b
}
