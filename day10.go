package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func init() {
	registerDay(10, day10)
}

type opcode10 int

const (
	noop = opcode10(0)
	addx = opcode10(1)
)

func (op opcode10) CycleLen() int {
	switch op {
	case noop:
		return 1
	case addx:
		return 2
	default:
		log.Fatalf("bad opcode %v", op)
		return -1
	}
}

type op10 struct {
	op opcode10
	n  int
}

func parse10(r io.Reader) ([]op10, error) {
	s := bufio.NewScanner(r)
	var dd []op10
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
		var n int
		if _, err := fmt.Sscanf(line, "noop"); err == nil {
			dd = append(dd, op10{op: noop})
		} else if _, err := fmt.Sscanf(line, "addx %d", &n); err == nil {
			dd = append(dd, op10{op: addx, n: n})
		} else {
			return nil, fmt.Errorf("failed to parse %q", line)
		}
	}
	return dd, s.Err()
}

func day10(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	dd, err := parse10(f)
	if err != nil {
		return nil, nil, err
	}
	cyc := 0
	x := 1

	part1 := 0
	screen := make([][40]byte, 6)

	for _, op := range dd {
		cl := op.op.CycleLen()
		cnLast := (cyc + 20) / 40
		cnNow := (cyc + cl + 20) / 40
		if cnLast != cnNow {
			cc := cnNow*40 - 20
			part1 += cc * x
		}
		for i := 0; i < cl; i++ {
			if abs(x-cyc%40) <= 1 {
				screen[cyc/40][cyc%40] = '#'
			} else {
				screen[cyc/40][cyc%40] = '.'
			}
			cyc += 1
		}
		if op.op == addx {
			x += op.n
		}
	}

	var part2 string
	for i, line := range screen {
		if i > 0 {
			part2 = part2 + "\n"
		}
		part2 += string(line[:])
	}

	return part1, part2, nil
}
