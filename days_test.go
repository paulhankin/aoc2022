package main

import (
	"fmt"
	"testing"
)

var wantFull [][2]any = [][2]any{
	1: {68787, 198041},
	2: {14264, 12382},
	3: {8401, 2641},
	4: {444, 801},
	5: {"WSFTMRHPP", "GSLCMFBRP"},
	6: {1802, 3551},
}

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

var wantPartial = []partialTest{
	{4, "day04_test.txt", 1, 2},
	{4, "day04_test.txt", 2, 4},
	{5, "day05_test.txt", 1, "CMZ"},
}

func TestDayPartial(t *testing.T) {
	for _, w := range wantPartial {
		t.Run(fmt.Sprintf("day%02d:%d on %s", w.day, w.part, w.filename), func(t *testing.T) {
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
