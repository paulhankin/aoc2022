package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func init() {
	registerDay(4, day04)
}

type day04d struct {
	a, b [2]int
}

func parse04(f io.Reader) ([]day04d, error) {
	s := bufio.NewScanner(f)
	var r []day04d
	for s.Scan() {
		t := strings.TrimSpace(s.Text())
		if t == "" {
			continue
		}
		var a, b, c, d int
		if _, err := fmt.Sscanf(t, "%d-%d,%d-%d", &a, &b, &c, &d); err != nil {
			return nil, err
		}
		r = append(r, day04d{[2]int{a, b}, [2]int{c, d}})
	}
	return r, s.Err()
}

func day04(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	dd, err := parse04(f)
	if err != nil {
		return nil, nil, err
	}
	N := len(dd)
	part1 := 0
	part2 := 0
	rc := func(a, b [2]int) bool {
		return b[1] <= a[1] && b[0] >= a[0]
	}
	rc2 := func(a, b [2]int) bool {
		return rc(a, b) || rc(b, a)
	}
	ol := func(a, b [2]int) bool {
		return !(a[0] > b[1] || b[0] > a[1])
	}
	ol2 := func(a, b [2]int) bool {
		return ol(a, b) || ol(b, a)
	}
	for i := 0; i < N; i++ {
		di := dd[i]
		if rc2(di.a, di.b) {
			part1++
		}
		if ol2(di.a, di.b) {
			part2++
		}
	}
	return part1, part2, nil
}
