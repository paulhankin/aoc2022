package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func parse14(r io.Reader) ([][][2]int, error) {
	s := bufio.NewScanner(r)
	var rr [][][2]int
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, " -> ")
		var track [][2]int
		for _, p := range parts {
			np := strings.Split(p, ",")
			if len(np) != 2 {
				return nil, fmt.Errorf("failed to parse a,b = %q", p)
			}
			a, err := strconv.Atoi(np[0])
			if err != nil {
				return nil, err
			}
			b, err := strconv.Atoi(np[1])
			if err != nil {
				return nil, err
			}
			track = append(track, [2]int{a, b})
		}
		rr = append(rr, track)
	}
	return rr, s.Err()
}

func day14(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	dd, err := parse14(f)
	if err != nil {
		return nil, nil, err
	}
	m := map[[2]int]rune{}
	maxy := 0
	for _, track := range dd {
		for i := 1; i < len(track); i++ {
			s, e := track[i-1], track[i]
			maxy = max(maxy, s[1])
			maxy = max(maxy, e[1])
			for j := 0; j <= max(abs(s[0]-e[0]), abs(s[1]-e[1])); j++ {
				x := s[0] + sgn(e[0]-s[0])*j
				y := s[1] + sgn(e[1]-s[1])*j
				m[[2]int{x, y}] = '#'
			}
		}
	}
	var parts []int
	for ppp := 0; ppp < 2; ppp++ {
		if ppp == 1 {
			for i := -10000; i < 10000; i++ {
				m[[2]int{i, maxy + 2}] = '#'
			}
		}
		for {
			x, y := 500, 0
			fallen := false
			blocked := false
			for {
				if m[[2]int{x, y + 1}] == 0 {
					y++
				} else if m[[2]int{x - 1, y + 1}] == 0 {
					y++
					x--
				} else if m[[2]int{x + 1, y + 1}] == 0 {
					y++
					x++
				} else {
					m[[2]int{x, y}] = 'o'
					if x == 500 && y == 0 {
						blocked = true
					}
					break
				}
				if y > 400 {
					fallen = true
					break
				}
			}
			if fallen {
				break
			}
			if blocked {
				break
			}
		}
		sand := 0
		for _, v := range m {
			if v == 'o' {
				sand++
			}
		}
		if false && strings.Contains(fn, "test") {
			for i := 0; i < 25; i++ {
				for j := 450; j <= 554; j++ {
					c := m[[2]int{j, i}]
					if c == 0 {
						fmt.Printf(" ")
					} else {
						fmt.Printf("%c", c)
					}
				}
				fmt.Println()
			}
		}
		parts = append(parts, sand)
	}
	return parts[0], parts[1], nil
}
