package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func parse19(r io.Reader) ([]byte, error) {
	s := bufio.NewScanner(r)
	var rr []byte
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		rr = append(rr, []byte(line)...)
	}
	return rr, s.Err()
}

func day19(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	dd, err := parse19(f)
	if err != nil {
		return nil, nil, err
	}
	_ = dd
	return 0, 0, nil
}
