package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	answer := part1()
	fmt.Println("Answer 1:", answer)
	answer = part2()
	fmt.Println("Answer 2:", answer)
}

func getLines() []string {
	content, err := ioutil.ReadFile("4/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(content), "\n")
}

func getAssignedZones(assignment string) []int {
	spl := strings.Split(assignment, "-")
	if len(spl) != 2 {
		panic("Invalid assignment")
	}
	start, err := strconv.Atoi(spl[0])
	if err != nil {
		panic("Invalid assignment")
	}
	end, err := strconv.Atoi(spl[1])
	if err != nil {
		panic("Invalid assignment")
	}
	zones := []int{}
	for i := start; i <= end; i++ {
		zones = append(zones, i)
	}
	return zones
}

func contains(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}

	return false
}

func doesAssignmentOverlap(line string) bool {
	assignments := strings.Split(line, ",")
	if len(assignments) != 2 {
		panic("Invalid line")
	}

	zones1 := getAssignedZones(assignments[0])
	zones2 := getAssignedZones(assignments[1])

	overlap := false 
	for _, zone1 := range zones1 {
		in := contains(zones2, zone1)
		if in {
			overlap = true
		} else {
			overlap = false
			break
		}
	}

	if overlap {
		return true
	}

	for _, zone2 := range zones2 {
		in := contains(zones1, zone2)
		if in {
			overlap = true
		} else {
			overlap = false
			break
		}
	}

	return overlap
}

func doesAssignmentOverlapAtAll(line string) bool {
	assignments := strings.Split(line, ",")
	if len(assignments) != 2 {
		panic("Invalid line")
	}

	zones1 := getAssignedZones(assignments[0])
	zones2 := getAssignedZones(assignments[1])

	for _, zone1 := range zones1 {
		in := contains(zones2, zone1)
		if in {
			return true
		}
	}

	for _, zone2 := range zones2 {
		in := contains(zones1, zone2)
		if in {
			return true
		}
	}

	return false
}

func part1() int {
	overlaps := 0
	lines := getLines()
	for _, line := range lines {
		overlap := doesAssignmentOverlap(line)
		if overlap {
			overlaps++
		}
	}
	return overlaps
}

func part2() int {
	overlaps := 0
	lines := getLines()
	for _, line := range lines {
		overlap := doesAssignmentOverlapAtAll(line)
		if overlap {
			overlaps++
		}
	}
	return overlaps
}