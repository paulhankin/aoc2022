package main

import (
	"flag"
	"fmt"
	"log"
)

type Day struct {
	File string
	F    func(string) (any, any, error)
}

var days = []Day{
	{},
	{"day01.txt", day01},
	{"day02.txt", day02},
	{"day03.txt", day03},
	{"day04.txt", day04},
	{"day05.txt", day05},
	{"day06.txt", day06},
}

var dayFlag = flag.Int("day", 0, "which day (0 = all)")

func main() {
	flag.Parse()

	if *dayFlag < 0 || *dayFlag >= len(days) {
		log.Fatalf("-day=%d out of range", *dayFlag)
	}

	for i := 1; i < len(days); i++ {
		if i == *dayFlag || *dayFlag == 0 {
			a, b, err := days[i].F(days[i].File)
			if err != nil {
				fmt.Printf("day%02d error: %v\n", i, err)
			}
			fmt.Printf("day%02d: %v %v\n", i, a, b)
		}
	}
}
