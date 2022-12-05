package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type move05 struct {
	n, from, to int
}

type data05 struct {
	crates [][]byte
	moves  []move05
}

func parse05(f io.Reader) (data05, error) {
	s := bufio.NewScanner(f)
	readCrates := false
	r := data05{
		crates: make([][]byte, 9),
	}
	for s.Scan() {
		if readCrates {
			line := strings.TrimSpace(s.Text())
			if line == "" {
				continue
			}
			var a, b, c int
			if _, err := fmt.Sscanf(line, "move %d from %d to %d", &a, &b, &c); err != nil {
				return r, err
			}
			r.moves = append(r.moves, move05{a, b, c})
		} else {
			line := s.Text()
			if strings.HasPrefix(strings.TrimSpace(line), "1") {
				readCrates = true
				continue
			}
			for i := 0; len(line) > 4*i; i++ {
				part := line[4*i : 4*i+3]
				if part[1] != ' ' {
					r.crates[i] = append(r.crates[i], part[1])
				}
			}
		}
	}
	for _, c := range r.crates {
		for i := 0; i < len(c)/2; i++ {
			c[i], c[len(c)-1-i] = c[len(c)-1-i], c[i]
		}
	}
	return r, s.Err()
}

func day05(filename string) (any, any, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	rs, err := parse05(f)
	if err != nil {
		return nil, nil, err
	}
	var parts []string
	for part := 0; part < 2; part++ {
		var crates [][]byte
		for _, c := range rs.crates {
			crates = append(crates, []byte(string(c)))
		}
		for _, m := range rs.moves {
			if part == 0 {
				for i := 0; i < m.n; i++ {
					fl := len(crates[m.from-1])
					got := crates[m.from-1][fl-1]
					crates[m.from-1] = crates[m.from-1][:fl-1]
					crates[m.to-1] = append(crates[m.to-1], got)
				}
			} else {
				fl := len(crates[m.from-1])
				got := crates[m.from-1][fl-m.n:]
				crates[m.from-1] = crates[m.from-1][:fl-m.n]
				crates[m.to-1] = append(crates[m.to-1], got...)
			}
		}
		var top []byte
		for _, c := range crates {
			if len(c) > 0 {
				top = append(top, c[len(c)-1])
			}
		}
		parts = append(parts, string(top))
	}
	return parts[0], parts[1], nil
}
