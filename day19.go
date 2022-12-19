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
	}
	return part1, 0, nil
}
