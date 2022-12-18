package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func init() {
	registerDay(15, day15)
}

type d15 struct {
	sensor [2]int
	beacon [2]int
}

func parse15(r io.Reader) ([]d15, error) {
	s := bufio.NewScanner(r)
	var rr []d15
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
		var x, y, bx, by int
		if _, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &x, &y, &bx, &by); err != nil {
			return nil, err
		}
		rr = append(rr, d15{sensor: [2]int{x, y}, beacon: [2]int{bx, by}})
	}
	return rr, s.Err()
}

func findRanges15(dd []d15, Y int) [][2]int {
	var changes [][2]int
	for _, d := range dd {
		dist := abs(d.sensor[0]-d.beacon[0]) + abs(d.sensor[1]-d.beacon[1])
		// find x0 and x1 such that distances from sensor to (x, Y) is dist.
		dy := abs(d.sensor[1] - Y)
		dx := dist - dy
		if dx < 0 {
			continue
		}
		x0 := d.sensor[0] - dx
		x1 := d.sensor[0] + dx
		changes = append(changes, [2]int{x0, 1})
		changes = append(changes, [2]int{x1 + 1, -1})
	}
	sort.Slice(changes, func(i, j int) bool {
		if changes[i][0] < changes[j][0] {
			return true
		}
		if changes[i][0] == changes[j][0] && changes[i][1] > changes[j][1] {
			return true
		}
		return false
	})

	var ranges [][2]int
	inside := 0
	start := 0
	for _, c := range changes {
		wasInside := (inside > 0)
		inside += c[1]
		nowInside := (inside > 0)
		if wasInside && !nowInside {
			ranges = append(ranges, [2]int{start, c[0]})
		} else if !wasInside && nowInside {
			start = c[0]
		}
	}
	return ranges
}

func day15(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	dd, err := parse15(f)
	if err != nil {
		return nil, nil, err
	}
	Y := 2000000
	isTest := strings.Contains(fn, "test")
	if isTest {
		Y = 10
	}

	bs := map[int]bool{}
	for _, d := range dd {
		if d.beacon[1] == Y {
			bs[d.beacon[0]] = true
		}
	}
	part1 := -len(bs)
	for _, r := range findRanges15(dd, Y) {
		part1 += r[1] - r[0]
	}
	ymax := 4000000
	if isTest {
		ymax = 20
	}
	part2 := func() int {
		for i := range dd {
			disti := abs(dd[i].sensor[0]-dd[i].beacon[0]) + abs(dd[i].sensor[1]-dd[i].beacon[1])
			for j := range dd {
				distj := abs(dd[j].sensor[0]-dd[j].beacon[0]) + abs(dd[j].sensor[1]-dd[j].beacon[1])
				for di := 0; di <= 1; di++ {
					for dj := 0; dj <= 1; dj++ {
						xpy := dd[i].sensor[0] + dd[i].sensor[1] - disti - di
						xmy := dd[j].sensor[0] - dd[j].sensor[1] + distj + dj
						if (xpy+xmy)%2 != 0 {
							continue
						}
						y := (xpy - xmy) / 2
						if 0 <= y && y <= ymax {
							r := findRanges15(dd, y)
							if len(r) == 2 {
								return 4000000*r[0][1] + y
							}
						}
					}
				}
			}
		}
		return -1
	}

	return part1, part2(), nil
}
