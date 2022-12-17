package main

import (
	"os"
)

func init() {
	registerDay(6, day06)
}

func day06(filename string) (any, any, error) {
	r, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}
	var parts [2]int
	for p := 0; p < 2; p++ {
		N := 4 + 10*p
		for i := range r {
			match := true
			for j := 0; match && j < N; j++ {
				for k := j + 1; match && k < N; k++ {
					if r[i+j] == r[i+k] {
						match = false
					}
				}
			}
			if match {
				parts[p] = i + N
				break
			}
		}
	}
	return parts[0], parts[1], nil
}
