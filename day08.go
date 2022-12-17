package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func init() {
	registerDay(8, day08)
}

func parse08(r io.Reader) ([]string, error) {
	s := bufio.NewScanner(r)
	var rr []string
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line != "" {
			rr = append(rr, line)
		}
	}
	return rr, s.Err()
}

func day08(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	dd, err := parse08(f)
	if err != nil {
		return nil, nil, err
	}
	N := len(dd)
	M := len(dd[0])
	vis := make([][]bool, N)
	for i := 0; i < N; i++ {
		vis[i] = make([]bool, M)
	}
	const noTree = byte(0)
	for i := 0; i < N; i++ {
		hi := noTree
		for j := 0; j < M; j++ {
			if dd[i][j] > hi {
				vis[i][j] = true
				hi = dd[i][j]
			}
		}
	}
	for i := 0; i < N; i++ {
		hi := noTree
		for j := M - 1; j >= 0; j-- {
			if dd[i][j] > hi {
				vis[i][j] = true
				hi = dd[i][j]
			}
		}
	}
	for j := 0; j < M; j++ {
		hi := noTree
		for i := 0; i < N; i++ {
			if dd[i][j] > hi {
				vis[i][j] = true
				hi = dd[i][j]
			}
		}
	}
	for j := 0; j < M; j++ {
		hi := noTree
		for i := N - 1; i >= 0; i-- {
			if dd[i][j] > hi {
				vis[i][j] = true
				hi = dd[i][j]
			}
		}
	}

	part1 := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if vis[i][j] {
				part1++
			}
		}
	}
	part2 := 0
	for i := 1; i < N-1; i++ {
		for j := 1; j < M-1; j++ {
			hi := dd[i][j]
			dirs := [][2]int{
				{1, 0}, {-1, 0}, {0, 1}, {0, -1},
			}
			var d [4]int
			for k := 0; k < 4; k++ {
				var dist int
				for dist = 1; true; dist++ {
					ii := i + dist*dirs[k][0]
					jj := j + dist*dirs[k][1]
					if ii < 0 || ii >= N || jj < 0 || jj >= M {
						dist--
						break
					}
					if dd[ii][jj] >= hi {
						break
					}
				}
				d[k] = dist
			}
			prod := d[0] * d[1] * d[2] * d[3]
			if prod > part2 {
				part2 = prod
			}
		}
	}

	return part1, part2, nil
}
