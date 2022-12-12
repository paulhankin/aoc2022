package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type d09 struct {
	c byte
	i int
}

func parse09(r io.Reader) ([]d09, error) {
	s := bufio.NewScanner(r)
	var rr []d09
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
		var d d09
		if _, err := fmt.Sscanf(line, "%c %d", &d.c, &d.i); err != nil {
			return nil, err
		}
		rr = append(rr, d)
	}
	return rr, s.Err()
}

func printChain(c [][2]int, n, m int) {
	for j := -n; j < n; j++ {
		for i := -m; i < m; i++ {
			r := '.'
			if i == 0 && j == 0 {
				r = 's'
			}
			for k := range c {
				if c[k][0] == i && c[k][1] == j {
					r = '0' + rune(len(c)-k-1)
				}
			}
			if r == '0' {
				r = 'H'
			}
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
	fmt.Println()
}

func day09(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	dd, err := parse09(f)
	if err != nil {
		return nil, nil, err
	}
	var parts [2]int
	for part := 1; part <= 2; part++ {
		pr := false // strings.Contains(fn, "test2") && part == 2
		N := 2 + 8*(part-1)
		chain := make([][2]int, N)
		if pr {
			printChain(chain, 20, 20)
		}
		visited := map[[2]int]bool{}
		for _, d := range dd {
			for n := 0; n < d.i; n++ {
				switch d.c {
				case 'U':
					chain[N-1][1] -= 1
				case 'D':
					chain[N-1][1] += 1
				case 'L':
					chain[N-1][0] -= 1
				case 'R':
					chain[N-1][0] += 1
				default:
					return nil, nil, fmt.Errorf("bad dir %c", d.c)
				}
				for i := N - 2; i >= 0; i-- {
					if abs(chain[i][0]-chain[i+1][0]) > 1 {
						chain[i][1] += sgn(chain[i+1][1] - chain[i][1])
						chain[i][0] += sgn(chain[i+1][0] - chain[i][0])
					} else if abs(chain[i][1]-chain[i+1][1]) > 1 {
						chain[i][0] += sgn(chain[i+1][0] - chain[i][0])
						chain[i][1] += sgn(chain[i+1][1] - chain[i][1])
					}
				}
				visited[chain[0]] = true
			}
			if pr {
				printChain(chain, 20, 20)
			}
		}
		parts[part-1] = len(visited)
	}
	return parts[0], parts[1], nil
}
