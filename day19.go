package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type mineral int

const (
	ore      = mineral(0)
	clay     = mineral(1)
	obsidian = mineral(2)
	geode    = mineral(3)
)

func (m mineral) String() string {
	switch m {
	case ore:
		return "ore"
	case clay:
		return "clay"
	case obsidian:
		return "obsidian"
	case geode:
		return "geode"
	}
	return "?"
}

func mineralFromString(s string) (mineral, bool) {
	for i := ore; i <= geode; i++ {
		if i.String() == s {
			return i, true
		}
	}
	return 0, false
}

type d19 struct {
	n     int
	robot [4][4]int
}

func init() {
	registerDay(19, day19)
}

func parse19(r io.Reader) ([]d19, error) {
	s := bufio.NewScanner(r)
	var rr []d19
	re, err := regexp.Compile("[0-9]+ [a-z]+")
	if err != nil {
		return nil, err
	}
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		line = strings.SplitN(line, ": ", 2)[1]
		robs := strings.Split(line, ".")
		entry := d19{
			n: len(rr) + 1,
		}
		if len(robs) != 5 {
			return nil, fmt.Errorf("wrong number of robots in %q", line)
		}
		for robIdx, r := range robs {
			r = strings.TrimSpace(r)
			if r == "" {
				continue
			}
			for _, ing := range re.FindAllString(r, -1) {
				var n int
				var min string
				if _, err := fmt.Sscanf(ing, "%d %s", &n, &min); err != nil {
					return nil, err
				}
				mm, ok := mineralFromString(min)
				if !ok {
					return nil, fmt.Errorf("unknown mineral %s", mm)
				}
				entry.robot[mineral(robIdx)][mm] = n
			}
		}
		rr = append(rr, entry)

	}
	return rr, s.Err()
}

type key19 struct {
	robots        [4]int
	ores          [4]int
	timeRemaining int
}

func genKeys(bp *d19, r int, oldk, k key19, res *[]key19) {
	if r == -1 {
		return
	}

	canBuild := true
	for i := 0; i < 4; i++ {
		if oldk.ores[i] < bp.robot[r][i] {
			canBuild = false
		}
	}
	if !(canBuild && (r > 1 || k.robots[r] == 0)) {
		genKeys(bp, r-1, oldk, k, res)
	}

	for k.robots[r] < 10 {
		k.robots[r]++
		for i := 0; i < 4; i++ {
			k.ores[i] -= bp.robot[r][i]
			oldk.ores[i] -= bp.robot[r][i]
			if oldk.ores[i] < 0 {
				return
			}
		}
		*res = append(*res, k)
		genKeys(bp, r-1, oldk, k, res)
	}
}

func geodes(bp *d19, key key19, cache map[key19]int) int {
	if i, ok := cache[key]; ok {
		return i
	}
	nk := key
	for i := 0; i < 4; i++ {
		nk.ores[i] += nk.robots[i]
	}
	nk.timeRemaining--
	if nk.timeRemaining == 0 {
		cache[key] = nk.ores[geode]
		return cache[key]
	}
	// fmt.Println(key)
	best := 0
	var keys []key19
	canBuildAll := true
	for r := 0; r < 4; r++ {
		for i := 0; i < 4; i++ {
			if key.ores[i] < bp.robot[r][i] {
				canBuildAll = false
			}
		}
	}
	if !canBuildAll {
		keys = append(keys, nk)
	}
	if nk.timeRemaining > 1 {
		// don't build robots on the last turn
		genKeys(bp, 3, key, nk, &keys)
	}
	for _, k := range keys {
		best = max(best, geodes(bp, k, cache))
	}
	cache[key] = best
	return best
}

func day19(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	dd, err := parse19(f)
	if err != nil {
		return nil, nil, err
	}
	part1 := 0
	for _, blueprint := range dd {
		fmt.Println(blueprint)
		cache := map[key19]int{}
		start := key19{
			robots:        [4]int{1, 0, 0, 0},
			timeRemaining: 24,
		}
		i := geodes(&blueprint, start, cache)
		fmt.Println(blueprint.n, i)
		part1 += i * blueprint.n
	}
	return part1, 0, nil
}
