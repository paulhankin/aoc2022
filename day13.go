package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"text/scanner"
)

type d13 struct {
	n  int
	ns []d13
}

func (d d13) String() string {
	if d.ns != nil {
		var parts []string
		for _, p := range d.ns {
			parts = append(parts, p.String())
		}
		return "[" + strings.Join(parts, ",") + "]"
	}
	return fmt.Sprintf("%d", d.n)
}

func parsed13(line string) (d13, error) {
	var sc scanner.Scanner
	sc.Init(strings.NewReader(line))
	var stack []d13
	stack = append(stack, d13{})
	for tok := sc.Scan(); tok != scanner.EOF; tok = sc.Scan() {
		if tok == '[' {
			stack = append(stack, d13{ns: []d13{}})
		} else if tok == ']' {
			top, ps := stack[len(stack)-1], stack[:len(stack)-1]
			stack = ps
			ps[len(ps)-1].ns = append(ps[len(ps)-1].ns, top)
		} else if tok == ',' {
		} else if tok == scanner.Int {
			n, err := strconv.Atoi(sc.TokenText())
			if err != nil {
				return d13{}, err
			}
			stack[len(stack)-1].ns = append(stack[len(stack)-1].ns, d13{n: n})
		} else {
			return d13{}, fmt.Errorf("failed to tokenize %c", tok)
		}
	}
	if len(stack) != 1 {
		return d13{}, fmt.Errorf("got %d stack frames after parsing: %v", len(stack), stack)
	}
	return stack[0].ns[0], nil
}

func parse13(r io.Reader) ([][2]d13, error) {
	s := bufio.NewScanner(r)
	var rr [][2]d13
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
		p0, err := parsed13(line)
		if err != nil {
			return nil, err
		}
		if !s.Scan() {
			break
		}
		p1, err := parsed13(strings.TrimSpace(s.Text()))
		if err != nil {
			return nil, err
		}
		rr = append(rr, [2]d13{p0, p1})
	}
	return rr, s.Err()
}

func compare13(a, b d13) int {
	if a.ns != nil && b.ns != nil {
		for i := 0; i < len(a.ns) && i < len(b.ns); i++ {
			if c := compare13(a.ns[i], b.ns[i]); c != 0 {
				return c
			}
		}
		if len(a.ns) < len(b.ns) {
			return -1
		} else if len(a.ns) > len(b.ns) {
			return 1
		} else {
			return 0
		}
	} else if a.ns == nil && b.ns == nil {
		if a.n < b.n {
			return -1
		}
		if a.n > b.n {
			return 1
		}
		return 0
	} else if a.ns == nil {
		return compare13(d13{ns: []d13{a}}, b)
	} else if b.ns == nil {
		return compare13(a, d13{ns: []d13{b}})
	}
	log.Fatalf("unreachable")
	return 0
}

func day13(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	dd, err := parse13(f)
	if err != nil {
		return nil, nil, err
	}
	part1 := 0
	var all []d13
	for i, d := range dd {
		all = append(all, d[0])
		all = append(all, d[1])
		if compare13(d[0], d[1]) < 0 {
			part1 += i + 1
		}
	}
	div2 := d13{ns: []d13{{ns: []d13{{n: 2}}}}}
	div6 := d13{ns: []d13{{ns: []d13{{n: 6}}}}}
	fmt.Println(div2, div6)

	all = append(all, div2, div6)
	sort.Slice(all, func(i, j int) bool {
		return compare13(all[i], all[j]) < 0
	})
	d2i, d6i := -1, -1
	for i, d := range all {
		if reflect.DeepEqual(d, div2) {
			d2i = i + 1
		}
		if reflect.DeepEqual(d, div6) {
			d6i = i + 1
		}
	}
	part2 := d2i * d6i

	return part1, part2, nil
}
