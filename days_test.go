package main

import (
	"flag"
	"fmt"
	"testing"
)

const day10part2 = `
###..####.#..#.####..##....##..##..###..
#..#....#.#..#.#....#..#....#.#..#.#..#.
#..#...#..####.###..#.......#.#....###..
###...#...#..#.#....#.##....#.#....#..#.
#.#..#....#..#.#....#..#.#..#.#..#.#..#.
#..#.####.#..#.#.....###..##...##..###..`

var wantFull [][2]any = [][2]any{
	1:  {68787, 198041},
	2:  {14264, 12382},
	3:  {8401, 2641},
	4:  {444, 801},
	5:  {"WSFTMRHPP", "GSLCMFBRP"},
	6:  {1802, 3551},
	7:  {1581595, 1544176},
	8:  {1705, 371200},
	9:  {6357, 2627},
	10: {13860, day10part2[1:]},
	11: {61503, 14081365540},
	12: {408, 399},
	13: {5393, 26712},
	14: {655, 26484},
	15: {4724228, 13622251246513},
	16: {1595, 2189},
	17: {3144, 1565242165201},
	18: {4636, 2572},
	19: {0, 0},
	20: {0, 0},
	21: {0, 0},
	22: {0, 0},
	23: {0, 0},
	24: {0, 0},
	25: {0, 0},
}

var skipSlow = flag.Bool("skip-slow", false, "skip slow tests")

var slowDays = map[int]bool{16: true}

func TestMissing(t *testing.T) {
	for i := range days {
		if i == 0 {
			continue
		}
		if i >= len(wantFull) || (wantFull[i][0] == nil && wantFull[i][1] == nil) {
			t.Errorf("missing tests for day %02d", i)
		}
	}
}

func TestDayFull(t *testing.T) {
	for i, want := range wantFull {
		if i == 0 {
			continue
		}
		t.Run(fmt.Sprintf("day%02d full", i), func(t *testing.T) {
			if *skipSlow && slowDays[i] {
				t.Skip("skipped day", i, "because of -skip-slow")
			}
			gota, gotb, goterr := days[i].F(days[i].File)
			if goterr != nil {
				t.Fatal(goterr)
			}
			if gota != want[0] || gotb != want[1] {
				t.Errorf("got %v:%T, %v:%T want %v:%T, %v:%T", gota, gota, gotb, gotb, want[0], want[0], want[1], want[1])
			}
		})
	}
}

type partialTest struct {
	day      int
	filename string
	part     int
	want     any
}

const day10test2p2 = `
##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`

var wantPartial = []partialTest{
	{4, "day04_test.txt", 1, 2},
	{4, "day04_test.txt", 2, 4},
	{5, "day05_test.txt", 1, "CMZ"},
	{7, "day07_test.txt", 1, 95437},
	{7, "day07_test.txt", 2, 24933642},
	{8, "day08_test.txt", 1, 21},
	{8, "day08_test.txt", 2, 8},
	{9, "day09_test.txt", 1, 13},
	{9, "day09_test.txt", 2, 1},
	{9, "day09_test2.txt", 2, 36},
	{10, "day10_test2.txt", 1, 13140},
	{10, "day10_test2.txt", 2, day10test2p2[1:]},
	{11, "day11_test.txt", 1, 10605},
	{11, "day11_test.txt", 2, 2713310158},
	{13, "day13_test.txt", 1, 13},
	{13, "day13_test.txt", 2, 140},
	{14, "day14_test.txt", 1, 24},
	{14, "day14_test.txt", 2, 93},
	{15, "day15_test.txt", 1, 26},
	{15, "day15_test.txt", 2, 56000011},
	{16, "day16_test.txt", 1, 1651},
	{16, "day16_test.txt", 2, 1707},
	{17, "day17_test.txt", 1, 3068},
	{17, "day17_test.txt", 2, 1514285714288},
	{18, "day18_test.txt", 1, 64},
	{18, "day18_test.txt", 2, 58},
	{19, "day19_test.txt", 1, 0},
	{19, "day19_test.txt", 2, 0},
	{20, "day20_test.txt", 1, 0},
	{20, "day20_test.txt", 2, 0},
	{21, "day21_test.txt", 1, 0},
	{21, "day21_test.txt", 2, 0},
	{22, "day22_test.txt", 1, 0},
	{22, "day22_test.txt", 2, 0},
	{23, "day23_test.txt", 1, 0},
	{23, "day23_test.txt", 2, 0},
	{24, "day24_test.txt", 1, 0},
	{24, "day24_test.txt", 2, 0},
	{25, "day25_test.txt", 1, 0},
	{25, "day25_test.txt", 2, 0},
}

func TestDayPartial(t *testing.T) {
	for _, w := range wantPartial {
		t.Run(fmt.Sprintf("day%02d:%d on %s", w.day, w.part, w.filename), func(t *testing.T) {
			if *skipSlow && slowDays[w.day] {
				t.Skip("skipped day", w.day, "because of -skip-slow")
			}
			gota, gotb, goterr := days[w.day].F(w.filename)
			if goterr != nil {
				t.Fatal(goterr)
			}
			if w.part < 1 || w.part > 2 {
				t.Fatalf("bad test case with part %d", w.part)
			}
			if w.part == 1 && gota != w.want {
				t.Errorf("got %v:%T want %v:%T", gota, gota, w.want, w.want)
			}
			if w.part == 2 && gotb != w.want {
				t.Errorf("got %v:%T want %v:%T", gotb, gotb, w.want, w.want)
			}
		})
	}
}
