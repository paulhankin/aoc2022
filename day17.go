package main

import (
	"bufio"
	"fmt"
	"hash/maphash"
	"io"
	"os"
	"strings"
)

func init() {
	registerDay(17, day17)
}

func parse17(r io.Reader) ([]byte, error) {
	s := bufio.NewScanner(r)
	var rr []byte
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		rr = append(rr, []byte(line)...)
	}
	return rr, s.Err()
}

var shapes = [][][2]int{
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
	{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}},
	{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
	{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
	{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
}

func move17(i, j, di, dj, R int, tower []byte) (int, int, bool) {
	for _, p := range shapes[R] {
		i1 := i + di + p[0]
		j1 := j + dj + p[1]
		if i1 < 0 || i1 >= 7 || j1 < 0 {
			return i, j, true
		}
		if tower[j1%TN]&(1<<i1) != 0 {
			return i, j, true
		}
	}
	return i + di, j + dj, false
}

func printTower(top int, tower [][7]byte) {
	for j := top + 2; j >= 0; j-- {
		fmt.Printf("|")
		for i := 0; i < 7; i++ {
			c := '.'
			if tower[j][i] != 0 {
				c = '#'
			}
			fmt.Printf("%c", c)
		}
		fmt.Printf("|\n")
	}
	fmt.Println()
}

const TN = 45

var hash17 maphash.Hash

func key17(tower []byte, top, R, gust int) uint64 {
	hash17.Reset()
	for i := 0; i < TN; i++ {
		hash17.WriteByte(tower[(top+i)%TN])
	}
	hash17.WriteByte(byte(top % TN))
	hash17.WriteByte(byte(R))
	hash17.WriteByte(byte(gust / 256))
	hash17.WriteByte(byte(gust % 256))
	return hash17.Sum64()
}

func day17(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	dd, err := parse17(f)
	if err != nil {
		return nil, nil, err
	}
	var res []int
	for parts := 0; parts < 2; parts++ {
		tower := make([]byte, TN)
		top := 0
		gust := -1
		ROCKN := 2022
		if parts == 1 {
			ROCKN = 1000000000000
		}
		found := map[uint64]int{}
		var tops []int
		for rock := 0; rock < ROCKN; rock++ {
			tops = append(tops, top)
			i, j := 2, top+3
			R := rock % 5
			for k := 0; k < 8; k++ {
				tower[(top+k)%TN] = 0
			}
			key := key17(tower, top, R, gust)
			if _, ok := found[key]; ok {
				from := found[key]
				to := rock

				// want to write ROCKN as from + K*(to - from) + k
				k := (ROCKN - from) % (to - from)
				K := (ROCKN - k - from) / (to - from)
				top = (tops[to]-tops[from])*K + tops[from+k]
				break
			}
			found[key] = rock
			for {
				gust = (gust + 1) % len(dd)
				dir := 0
				if dd[gust] == '<' {
					dir = -1
				} else if dd[gust] == '>' {
					dir = 1
				} else {
					return nil, nil, fmt.Errorf("bad dir")
				}

				i, j, _ = move17(i, j, dir, 0, R, tower)
				var stuck bool
				i, j, stuck = move17(i, j, 0, -1, R, tower)
				if stuck {
					break
				}
			}
			for _, p := range shapes[R] {
				tower[(j+p[1])%TN] |= 1 << (i + p[0])
				top = max(top, j+p[1]+1)
			}
		}
		res = append(res, top)
	}
	return res[0], res[1], nil
}
