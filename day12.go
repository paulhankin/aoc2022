package main

import (
	"bytes"
	"os"
)

func day12(fn string) (any, any, error) {
	dd, err := os.ReadFile(fn)
	if err != nil {
		return nil, nil, err
	}
	D := bytes.Split(dd, []byte("\n"))

	N := len(D)
	M := len(D[0])
	var start, end [2]int
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if D[i][j] == 'S' {
				start = [2]int{i, j}
			}
			if D[i][j] == 'E' {
				end = [2]int{i, j}
			}
		}
	}
	const MAX_INT = 1000000000
	dist := make(map[[2]int]int)
	var Q [][2]int
	Q = append(Q, end)
	dist[end] = 0
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	H := func(x [2]int) int {
		if x[0] < 0 || x[0] >= N || x[1] < 0 || x[1] >= M {
			return MAX_INT
		}
		c := D[x[0]][x[1]]
		if c == 'S' {
			c = 'a'
		} else if c == 'E' {
			c = 'z'
		}
		return int(c) - 'a'
	}
	for len(Q) > 0 {
		var x [2]int
		x, Q = Q[0], Q[1:]
		for _, d := range dirs {
			x1 := [2]int{x[0] + d[0], x[1] + d[1]}
			if _, ok := dist[x1]; ok {
				continue
			}
			h0 := H(x)
			h1 := H(x1)
			if h1 < h0-1 || h1 == MAX_INT {
				continue
			}
			dist[x1] = dist[x] + 1
			Q = append(Q, x1)
		}
	}

	shortest := MAX_INT
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			s := [2]int{i, j}
			d, ok := dist[s]
			if H(s) == 0 && ok {
				shortest = min(shortest, d)
			}
		}
	}

	return dist[start], shortest, nil

}
