package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func parse03(r io.Reader) ([][2]string, error) {
	var res [][2]string
	s := bufio.NewScanner(r)
	for s.Scan() {
		t := strings.TrimSpace(s.Text())
		if t == "" {
			continue
		}
		a, b := t[:len(t)/2], t[len(t)/2:]
		if len(a) != len(b) {
			return nil, fmt.Errorf("%q and %q different lengths", a, b)
		}
		res = append(res, [2]string{a, b})
	}
	return res, s.Err()
}

func findShared(a, b string) []byte {
	var got []byte
	for _, c := range a {
		if bytes.ContainsRune(got, rune(c)) {
			continue
		}
		if strings.ContainsRune(b, rune(c)) {
			got = append(got, byte(c))
		}
	}
	return got
}

func day03(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	rs, err := parse03(f)
	if err != nil {
		return nil, nil, err
	}

	pri := func(i byte) int {
		if i >= 'a' && i <= 'z' {
			return int(i - 'a' + 1)
		} else if i >= 'A' && i <= 'Z' {
			return int(i - 'A' + 27)
		} else {
			log.Fatal("bad")
			return 0
		}
	}

	part1 := 0
	for _, p := range rs {
		shared := findShared(p[0], p[1])
		if len(shared) != 1 {
			return nil, nil, fmt.Errorf("not exactly one shared item in %q and %q", p[0], p[1])
		}
		part1 += pri(shared[0])
	}

	J := func(s [2]string) string {
		return s[0] + s[1]
	}

	part2 := 0
	for i := 0; i < len(rs); i += 3 {
		s := string(findShared(J(rs[i]), J(rs[i+1])))
		shared := findShared(s, J(rs[i+2]))
		if len(shared) != 1 {
			return nil, nil, fmt.Errorf("not exactly one shared item in %q and %q and %q", rs[i], rs[i+1], rs[i+2])
		}
		part2 += pri(shared[0])

	}

	return part1, part2, nil
}
