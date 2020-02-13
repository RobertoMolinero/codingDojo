package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	names := ReadTestData("../names.txt")
	sort.Strings(names)

	sum := GetNameScore(names)
	fmt.Printf("%d\n", sum)
}

func ReadTestData(filename string) []string {
	var names []string

	dat, _ := ioutil.ReadFile(filename)
	r := csv.NewReader(strings.NewReader(string(dat)))
	r.Comma = ','

	records, _ := r.Read()

	for _, v := range records {
		names = append(names, v)
	}

	return names
}

func GetNameScore(names []string) int {
	var sum int

	for i, name := range names {
		alphabeticalValue := GetAlphabeticalValueOfName(name)
		sortValue := i + 1
		score := alphabeticalValue * sortValue
		sum += score
	}

	return sum
}

func GetAlphabeticalValueOfName(name string) int {
	var r = []rune(name)
	var sum int

	for _, v := range r {
		sum += int(v) - 64
	}

	return sum
}
