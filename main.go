package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)


func part1() int {
    content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	elfs := make(map[int]int)

	i := 1
	elfs[i] = 0
	for _, line := range lines {
		if line == "" {
			i++
			elfs[i] = 0
			continue
		}
		cal, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		elfs[i] += cal
	}

	highestCal := 0
	for _, cal := range elfs {
		if cal > highestCal {
			highestCal = cal
		}
	}

	return highestCal
}

// mostly redone p1
func part2() int {
    content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	elfs := []int{0}

	i := 0
	for _, line := range lines {
		if line == "" {
			i++
			elfs = append(elfs, 0)
			continue
		}
		cal, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		elfs[i] += cal
	}

    sort.Sort(sort.Reverse(sort.IntSlice(elfs)))

	first := elfs[0]
	second := elfs[1]
	third := elfs[2]

	return first + second + third
}

func main() {
	answer := part1()
	fmt.Println("Answer 1:", answer)
	answer = part2()
	fmt.Println("Answer 2:", answer)
}