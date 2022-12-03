package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

func main() {
	answer := part1()
	fmt.Println("Answer 1:", answer)
	answer = part2()
	fmt.Println("Answer 2:", answer)
}

func getCompartments(line string) (string, string){
	comparmentSize := len(line) / 2
	return line[0:comparmentSize], line[comparmentSize:]
}

func findDups(comp1 string, comp2 string) string {
	for _, char := range comp1 {
		if strings.Contains(comp2, string(char)) {
			return string(char)

		}
	}
	panic("No dups found")
}

func getDupValue(dup rune) int {
	lowDup := strings.ToLower(string(dup))
	val := 0
	switch lowDup {
	case "a":
		val =  1
	case "b":
		val =  2
	case "c":
		val =  3
	case "d":
		val =  4
	case "e":
		val =  5
	case "f":
		val =  6
	case "g":
		val =  7
	case "h":
		val =  8
	case "i":
		val =  9
	case "j":
		val =  10
	case "k":
		val =  11
	case "l":
		val =  12
	case "m":
		val =  13
	case "n":
		val =  14
	case "o":
		val =  15
	case "p":
		val =  16
	case "q":
		val =  17
	case "r":
		val =  18
	case "s":
		val =  19
	case "t":
		val =  20
	case "u":
		val =  21
	case "v":
		val =  22
	case "w":
		val =  23
	case "x":
		val =  24
	case "y":
		val =  25
	case "z":
		val =  26
	default:
		panic("Unknown char")
	}
	if unicode.IsUpper(rune(dup)) {
		return val + 26
	}
	return val
}


func part1() int {
	content, err := ioutil.ReadFile("3/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	total := 0
	for _, line := range lines {
		comp1, comp2 := getCompartments(line)
		dup := findDups(comp1, comp2)
		dupVal := getDupValue(rune(dup[0]))
		total += dupVal
	}
	return total
}

func compareComartments(elf1 string, elf2 string, elf3 string) string {
	for _, char := range elf1 {
		if strings.Contains(elf2, string(char)) && strings.Contains(elf3, string(char)) {
			return string(char)
		}
	}
	panic("No dups found")
}

func part2() int {
	content, err := ioutil.ReadFile("3/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	total := 0
	elfLines := []string{}
	for _, line := range lines {
		elfLines = append(elfLines, line)
		if len(elfLines) == 3 {
			dup := compareComartments(elfLines[0], elfLines[1], elfLines[2])
			dupVal := getDupValue(rune(dup[0]))
			total += dupVal
			elfLines = []string{}
		}
	}
	return total
}