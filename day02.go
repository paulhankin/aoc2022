package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func parse02(r io.Reader) ([][2]rune, error) {
	s := bufio.NewScanner(r)
	var res [][2]rune
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
		var a, b rune
		if _, err := fmt.Sscanf(line, "%c %c", &a, &b); err != nil {
			return nil, err
		}
		res = append(res, [2]rune{a, b})
	}
	return res, s.Err()
}

func score02(i, j int) int {
	w := (j - i + 3) % 3
	if w == 1 {
		return 6 + j
	} else if w == 0 {
		return 3 + j
	} else {
		return 0 + j
	}
}

func day02(f string) (any, any, error) {
	var strat [][2]rune
	if err := func() error {
		r, err := os.Open(f)
		if err != nil {
			return err
		}
		defer r.Close()
		strat, err = parse02(r)
		return err
	}(); err != nil {
		return nil, nil, err
	}
	mapa := map[rune]int{
		'A': 1,
		'B': 2,
		'C': 3,
	}
	mapb := map[rune]int{
		'X': 1,
		'Y': 2,
		'Z': 3,
	}

	score1, score2 := 0, 0

	for _, p := range strat {
		a := mapa[p[0]]
		score1 += score02(a, mapb[p[1]])
		var b int
		if p[1] == 'X' {
			b = (a-1+2)%3 + 1
		} else if p[1] == 'Y' {
			b = a
		} else if p[1] == 'Z' {
			b = a%3 + 1
		} else {
			log.Fatal(p[1])
		}
		score2 += score02(a, b)
	}

	return score1, score2, nil

}
