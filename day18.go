package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func init() {
	registerDay(18, day18)
}

func parse18(r io.Reader) (map[[3]int]bool, error) {
	s := bufio.NewScanner(r)
	rr := map[[3]int]bool{}
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		var x, y, z int
		if _, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z); err != nil {
			return nil, err
		}
		rr[[3]int{x, y, z}] = true
	}
	return rr, s.Err()
}

var dirs3d = [][3]int{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}

func day18(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	dd, err := parse18(f)
	if err != nil {
		return nil, nil, err
	}
	part1 := 0
	minc := [3]int{100, 100, 100}
	maxc := [3]int{-100, -100, -100}
	for pos := range dd {
		for i := 0; i < 3; i++ {
			minc[i] = min(minc[i], pos[i])
			maxc[i] = max(maxc[i], pos[i])
		}
		for _, dir := range dirs3d {
			p2 := [3]int{pos[0] + dir[0], pos[1] + dir[1], pos[2] + dir[2]}
			if _, ok := dd[p2]; !ok {
				part1++
			}
		}
	}
	var airs [][3]int
	airind := map[[3]int]int{}
	for x := minc[0] - 1; x <= maxc[0]+1; x++ {
		for y := minc[1] - 1; y <= maxc[1]+1; y++ {
			for z := minc[2] - 1; z <= maxc[2]+1; z++ {
				if _, ok := dd[[3]int{x, y, z}]; !ok {
					airind[[3]int{x, y, z}] = len(airs)
					airs = append(airs, [3]int{x, y, z})
				}
			}
		}
	}
	uf := NewUnionFind(len(airs))
	for i, a := range airs {
		for _, dir := range dirs3d {
			p2 := [3]int{a[0] + dir[0], a[1] + dir[1], a[2] + dir[2]}
			if _, ok := dd[p2]; ok {
				continue
			}
			ni, ok := airind[p2]
			if !ok {
				continue // out of range, or not an air
			}
			uf.Union(i, ni)
		}
	}
	af := uf.Find(airind[[3]int{minc[0] - 1, minc[1] - 1, minc[2] - 1}])

	var part2 int
	for pos := range dd {
		for _, dir := range dirs3d {
			p2 := [3]int{pos[0] + dir[0], pos[1] + dir[1], pos[2] + dir[2]}
			ni, ok := airind[p2]
			if !ok {
				continue
			}
			if uf.Find(ni) == af {
				part2++
			}
		}
	}

	return part1, part2, nil
}
