package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

type d16s struct {
	name    string
	flow    int
	leadsTo []string
}

type targetDist struct {
	target int
	dist   int
}

type d16 struct {
	flows []int
	dists [][]int
}

type nd16 struct {
	name string
	dist int
}

func shortestPath16(nodes []d16s, names map[string]int, from, to string) int {
	var Q []nd16
	Q = append(Q, nd16{from, 0})
	visited := map[string]bool{}
	for len(Q) > 0 {
		var q nd16
		q, Q = Q[0], Q[1:]
		if q.name == to {
			return q.dist
		}
		for _, lt := range nodes[names[q.name]].leadsTo {
			if visited[lt] {
				continue
			}
			visited[lt] = true
			Q = append(Q, nd16{lt, q.dist + 1})
		}
	}
	log.Fatalf("failed to find path from %q to %q", from, to)
	return 0
}

func parse16(r io.Reader) (d16, int, error) {
	s := bufio.NewScanner(r)
	var rr []d16s
	names := map[string]int{}
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
		var name string
		var rate int
		if _, err := fmt.Sscanf(line, "Valve %s has flow rate=%d;", &name, &rate); err != nil {
			return d16{}, 0, fmt.Errorf("%q: %v", line, err)
		}
		parts := strings.Split(line, "to valve")
		if len(parts) != 2 {
			return d16{}, 0, fmt.Errorf("couldn't find valves in line: %q", line)
		}
		parts[1] = strings.Trim(parts[1], "s ")
		parts = strings.Split(parts[1], ", ")
		for _, p := range parts {
			if len(p) != 2 {
				return d16{}, 0, fmt.Errorf("bad leadsto %q", p)
			}
		}
		names[name] = len(rr)
		rr = append(rr, d16s{name: name, flow: rate, leadsTo: parts})
	}
	cnames := []string{}
	flows := []int{}
	dists := [][]int{}
	AA := -1
	for _, from := range rr {
		if from.name != "AA" && from.flow == 0 {
			continue
		}
		if from.name == "AA" {
			AA = len(cnames)
		}
		cnames = append(cnames, from.name)
		flows = append(flows, from.flow)
	}
	for _, from := range cnames {
		var distRow []int
		for _, to := range cnames {
			distRow = append(distRow, shortestPath16(rr, names, from, to))
		}
		dists = append(dists, distRow)
	}
	return d16{
		flows: flows,
		dists: dists,
	}, AA, s.Err()

}

type s16 struct {
	idx    int
	opened uint64
	time   int
}

func search16(dd d16, mask uint64, idx int, opened uint64, time int, cache map[s16]int) int {
	if time <= 1 {
		return 0
	}
	key := s16{idx, opened, time}
	if v, ok := cache[key]; ok {
		return v
	}
	f := 0
	if time > 1 && (opened>>idx)&1 == 0 {
		ff := (time-1)*dd.flows[idx] + search16(dd, mask, idx, opened|(uint64(1)<<idx), time-1, cache)
		f = max(f, ff)
	}
	for j, dist := range dd.dists[idx] {
		if idx == j || time <= dist+1 || (mask>>j)&1 == 0 {
			continue
		}
		ff := search16(dd, mask, j, opened, time-dist, cache)
		f = max(f, ff)
	}
	cache[key] = f
	return f
}

func makeSubsets(n, aa int) []uint64 {
	if n == 0 {
		return []uint64{0}
	}
	s := makeSubsets(n-1, aa)
	if n == aa+1 {
		return s
	}
	var s2 []uint64
	for _, x := range s {
		s2 = append(s2, x|(uint64(1)<<(n-1)))
	}
	return append(s, s2...)
}

func day16(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	dd, AA, err := parse16(f)
	if err != nil {
		return nil, nil, err
	}
	part1 := search16(dd, (1<<62)-1, AA, 0, 30, map[s16]int{})

	subsets := makeSubsets(len(dd.flows), AA)
	solns := map[uint64]int{}
	var mut sync.Mutex
	work := make(chan uint64, 16)
	go func() {
		for _, s := range subsets {
			work <- s
		}
		close(work)
	}()

	var wg sync.WaitGroup
	for workers := 0; workers < 16; workers++ {
		wg.Add(1)
		go func() {
			for s := range work {
				r := search16(dd, s|(uint64(1)<<AA), AA, 0, 26, map[s16]int{})
				mut.Lock()
				solns[s] = r
				mut.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	part2 := 0
	mask := (uint64(1) << len(dd.flows)) - 1
	mask &^= (uint64(1) << AA)
	for s, f := range solns {
		si := s ^ mask
		f2, ok := solns[si]
		if !ok {
			log.Fatalf("failed to find complementary subset %x -> %x", s, si)
		}
		part2 = max(part2, f+f2)
	}

	return part1, part2, err
}
