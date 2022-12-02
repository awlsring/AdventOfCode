package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	answer := part1()
	fmt.Println("Answer 1:", answer)
	answer = part2()
	fmt.Println("Answer 2:", answer)
}

type Choice struct {
	cType string
	cPoint int
}

func buildChoice(choice string) Choice {
	switch true {
	case (choice == "A" || choice == "X"):
		return Choice{"rock", 1}
	case (choice == "B" || choice == "Y"):
		return Choice{"paper", 2}
	case (choice == "C" || choice == "Z"):
		return Choice{"scissors", 3}
	}
	panic("Invalid choice")
}

func decideRound(elf string, descision string) int {
	result := 0
	switch descision {
	case "Y":
		result = 3
	case "Z":
		result = 6
	}

	switch true {
	case (elf == "A" && descision == "X"):
		fmt.Println("Elf played rock, I want to lose and play scissors")
		return 3 + result
	case (elf == "B" && descision == "X"):
		fmt.Println("Elf played paper, I want to lose and play rock")
		return 1 + result
	case (elf == "C" && descision == "X"):
		fmt.Println("Elf played scissors, I want to lose and play paper")
		return 2 + result
	case (elf == "A" && descision == "Y"):
		fmt.Println("Elf played rock, I want to tie and play rock")
		return 1 + result
	case (elf == "B" && descision == "Y"):
		fmt.Println("Elf played paper, I want to tie and play paper")
		return 2 + result
	case (elf == "C" && descision == "Y"):
		fmt.Println("Elf played scissors, I want to tie and play scissors")
		return 3 + result
	case (elf == "A" && descision == "Z"):
		fmt.Println("Elf played rock, I want to win and play paper")
		return 2 + result
	case (elf == "B" && descision == "Z"):
		fmt.Println("Elf played paper, I want to win and play scissors")
		return 3 + result
	case (elf == "C" && descision == "Z"):
		fmt.Println("Elf played scissors, I want to win and play rock")
		return 1 + result
	}
	panic("Invalid choice")
}

func determinePoints(you int, me int) int {
	if you == me {
		return 3 + me
	}
	if you == 1 && me == 3 {
		return 0 + me
	}
	if you == 3 && me == 1 {
		return 6 + me
	}
	if you > me {
		return 0 + me
	}
	if you < me {
		return 6 + me
	}
	panic("Invalid winner")
}

func parseLine(line string) (string, string) {
	ans := strings.Split(line, " ")
	if len(ans) != 2 {
		panic("Invalid line")
	}
	return ans[0], ans[1]
}

func part1() int {
	content, err := ioutil.ReadFile("2/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	total := 0
	for _, line := range lines {
		you, me := parseLine(line)
		youChoice := buildChoice(you)
		meChoice := buildChoice(me)
		roundPoints := determinePoints(youChoice.cPoint, meChoice.cPoint)
		total += roundPoints
		fmt.Println("Elf played", youChoice.cType, "I play", meChoice.cType)
		fmt.Println("I get", roundPoints, "points")
	}

	return total
}

func part2() int {
	content, err := ioutil.ReadFile("2/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	total := 0
	for _, line := range lines {
		you, me := parseLine(line)
		roundPoints := decideRound(you, me)
		total += roundPoints
		fmt.Println("I get", roundPoints, "points")
	}

	return total
}