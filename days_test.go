package main

import (
	"fmt"
	"testing"
)

var wantFull [][2]string = [][2]string{
	1: {"68787", "198041"},
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
				t.Errorf("got %q,%q want %q,%q", gota, gotb, want[0], want[1])
			}
		})
	}
}
