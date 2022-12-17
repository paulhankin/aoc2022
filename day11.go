package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func init() {
	registerDay(11, day11)
}

type data11 struct {
	monkey              int
	startItems          []int
	opCode              rune
	opArg               int
	divBy               int
	whenTrue, whenFalse int
}

func parse11(r io.Reader) ([]data11, error) {
	var d11 []data11
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
		var n int
		var op rune
		if _, err := fmt.Sscanf(line, "Monkey %d", &n); err == nil {
			d11 = append(d11, data11{monkey: n})
		} else if _, err := fmt.Sscanf(line, "Starting items:"); err == nil {
			p1 := strings.SplitN(line, ":", 2)
			parts := strings.Split(p1[1], ", ")
			for _, p := range parts {
				p = strings.TrimSpace(p)
				if p == "" {
					continue
				}
				pn, err := strconv.Atoi(p)
				if err != nil {
					return nil, err
				}
				d11[len(d11)-1].startItems = append(d11[len(d11)-1].startItems, pn)
			}
		} else if _, err := fmt.Sscanf(line, "Operation: new = old %c %d", &op, &n); err == nil {
			d11[len(d11)-1].opCode = op
			d11[len(d11)-1].opArg = n
		} else if _, err := fmt.Sscanf(line, "Operation: new = old * old"); err == nil {
			d11[len(d11)-1].opCode = 'S'
		} else if _, err := fmt.Sscanf(line, "Test: divisible by %d", &n); err == nil {
			d11[len(d11)-1].divBy = n
		} else if _, err := fmt.Sscanf(line, "If true: throw to monkey %d", &n); err == nil {
			d11[len(d11)-1].whenTrue = n
		} else if _, err := fmt.Sscanf(line, "If false: throw to monkey %d", &n); err == nil {
			d11[len(d11)-1].whenFalse = n
		} else {
			return nil, fmt.Errorf("line %q not parsed", line)
		}
	}
	return d11, s.Err()
}

func day11(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	d11, err := parse11(f)
	if err != nil {
		return nil, nil, err
	}

	var partscore []int
	for part := 1; part <= 2; part++ {
		items := make([][]int, len(d11))
		for i, d := range d11 {
			for _, it := range d.startItems {
				items[i] = append(items[i], it)
			}
		}
		icount := make([]int, len(d11))
		rt := 20
		if part == 2 {
			rt = 10000
		}
		M := 2 * 3 * 5 * 7 * 11 * 13 * 17 * 19 * 23
		for round := 0; round < rt; round++ {
			for m := 0; m < len(d11); m++ {
				md := d11[m]
				icount[m] += len(items[m])
				for _, it := range items[m] {
					if md.opCode == '+' {
						it = it + md.opArg
					} else if md.opCode == '*' {
						it = it * md.opArg
					} else if md.opCode == 'S' {
						it = it * it
					} else {
						log.Fatalf("unknown opcode %c", md.opCode)
					}
					if part == 1 {
						it = it / 3
					} else {
						it = it % M
					}
					dt := it % md.divBy
					if dt == 0 {
						items[md.whenTrue] = append(items[md.whenTrue], it)
					} else {
						items[md.whenFalse] = append(items[md.whenFalse], it)
					}
				}
				items[m] = nil
			}
		}

		sort.Slice(icount, func(i, j int) bool {
			return icount[i] > icount[j]
		})
		partscore = append(partscore, icount[0]*icount[1])
	}

	return partscore[0], partscore[1], nil
}
