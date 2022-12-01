package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"strings"
	"sort"
	"strconv"
)

func parseInput(r io.Reader) ([]int, error) {
	s := bufio.NewScanner(r)
	var res []int
	need := true
	for s.Scan() {
		t := strings.TrimSpace(s.Text())
		if len(t) == 0 {
			need = true
			continue
		}
		i, err := strconv.ParseInt(t, 10, 0)
		if err != nil {
			return nil, err
		}
		if need {
			res = append(res, 0)
			need = false
		}
		res[len(res)-1] += int(i)
	}
	return res, s.Err()
}

func day01() (int, int, error) {
	f, err := os.Open("day01.txt")
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()
	cals, err := parseInput(f)
	if err != nil {
		return 0, 0, err
	}

	sort.Slice(cals, func(i, j int) bool {
		return cals[i] > cals[j]
	})

	return cals[0], cals[0] + cals[1] + cals[2], nil

}

func main() {
	a, b, err := day01()
	if err != nil {
		panic(err)
	}
	fmt.Println(a, b)
}