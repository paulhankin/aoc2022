package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func init() {
	registerDay(22, day22)
}

func parse22(r io.Reader) ([]byte, error) {
	s := bufio.NewScanner(r)
	var rr []byte
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		rr = append(rr, []byte(line)...)
	}
	return rr, s.Err()
}

func day22(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	dd, err := parse22(f)
	if err != nil {
		return nil, nil, err
	}
	_ = dd
	return 0, 0, nil
}
