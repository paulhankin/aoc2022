package main

import (
	"fmt"
	"testing"
)

var wantFull [][2]any = [][2]any{
	1: {68787, 198041},
	2: {14264, 12382},
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
