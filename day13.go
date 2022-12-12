package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type d13 struct {
}

func parse13(r io.Reader) (d13, error) {
	s := bufio.NewScanner(r)
	var rr d13
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
	}
	return rr, s.Err()
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
	fmt.Println(dd)

	return 0, 0, nil
}
