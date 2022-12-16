package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type d17 struct{}

func parse17(r io.Reader) (d17, error) {
	s := bufio.NewScanner(r)
	var rr d17
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
	}
	return rr, s.Err()
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
	fmt.Println(dd)
	return 0, 0, nil
}
