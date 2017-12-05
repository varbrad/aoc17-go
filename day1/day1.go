package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func solve1(n string) int {
	sum := 0
	for i := 0; i < len(n)-1; i++ {
		if n[i] == n[i+1] {
			v, err := strconv.Atoi(string(n[i]))
			check(err)
			sum += v
		}
	}
	if n[0] == n[len(n)-1] {
		v, err := strconv.Atoi(string(n[0]))
		check(err)
		sum += v
	}
	return sum
}

func solve2(n string) int {
	l := len(n) / 2
	sum := 0
	for i := 0; i < l; i++ {
		if n[i] == n[i+l] {
			v, err := strconv.Atoi(string(n[i]))
			check(err)
			sum += v
		}
	}
	return sum * 2
}

func main() {
	data, err := ioutil.ReadFile("./day1/input.txt")
	check(err)
	fmt.Println("Day 1, Part 1:", solve1(string(data)))
	fmt.Println("Day 1, Part 2:", solve2(string(data)))
}
