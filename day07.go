package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func init() {
	registerDay(7, day07)
}

type cmd07 struct {
	command string
	output  []string
}

type fileSize struct {
	name string
	size int
}

type dir07 struct {
	name  string
	size  int
	dirs  []string
	files []fileSize
}

func parse07(f io.Reader) ([]cmd07, error) {
	s := bufio.NewScanner(f)
	var r []cmd07
	for s.Scan() {
		t := strings.TrimSpace(s.Text())
		if t == "" {
			continue
		}
		if strings.HasPrefix(t, "$ ") {
			r = append(r, cmd07{
				command: t[2:],
			})
		} else {
			r[len(r)-1].output = append(r[len(r)-1].output, t)
		}
	}
	return r, s.Err()
}

func day07(fn string) (any, any, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	dd, err := parse07(f)
	if err != nil {
		return nil, nil, err
	}

	dirs := map[string]*dir07{}
	var cwd []string
	cwds := func(dirs []string) string {
		return strings.Join(dirs, "/")
	}
	for _, cmd := range dd {
		var d string
		if _, err := fmt.Sscanf(cmd.command, "cd %s", &d); err == nil {
			if d == ".." {
				cwd = cwd[:len(cwd)-1]
			} else if d == "/" {
				cwd = nil
			} else if d == "" {
				log.Fatal("empty cd")
			} else {
				cwd = append(cwd, d)
			}
		} else if _, err := fmt.Sscanf(cmd.command, "ls"); err == nil {
			if _, ok := dirs[cwds(cwd)]; !ok {
				dirs[cwds(cwd)] = &dir07{
					name: cwds(cwd),
					size: -1,
				}
			}
			d07 := dirs[cwds(cwd)]
			for _, line := range cmd.output {
				var name string
				var size int
				if _, err := fmt.Sscanf(line, "dir %s", &name); err == nil {
					d07.dirs = append(d07.dirs, cwds(append(cwd, name)))
				} else if _, err := fmt.Sscanf(line, "%d %s", &size, &name); err == nil {
					d07.files = append(d07.files, fileSize{name, size})
				} else {
					log.Fatalf("failed to parse command output line %q", line)
				}
			}
		}
	}
	var getSize func(d string) int
	getSize = func(d string) int {
		dir, ok := dirs[d]
		if !ok {
			log.Fatalf("failed to find dir %q", d)
		}
		if dir.size != -1 {
			return dir.size
		}
		dir.size = 0
		for _, f := range dir.files {
			dir.size += f.size
		}
		for _, sd := range dir.dirs {
			dir.size += getSize(sd)
		}
		return dir.size
	}
	var ds []fileSize
	for d := range dirs {
		ds = append(ds, fileSize{d, getSize(d)})
	}
	var part1 int
	part2 := dirs[""].size
	ts := 70000000 - dirs[""].size
	need := 30000000 - ts
	for _, d := range ds {
		if d.size <= 100000 {
			part1 += d.size
		}
		if d.size >= need && d.size < part2 {
			part2 = d.size
		}
	}
	return part1, part2, nil
}
