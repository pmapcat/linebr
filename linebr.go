package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const NEWLINE = "\n"

const BUFLEN = 4 * 4096

type Proc int

const (
	PROC_TEXT Proc = iota
	PROC_DEFAULTLEN
	PROC_NOOP
	PROC_REGEX
	PROC_FUNC
)

func pagedStdin(pattern string, proc Proc, buflen int) {
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, buflen)
	rrep := pattern + NEWLINE

	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			panic(err)
		}

		switch {
		case proc == PROC_TEXT:
			fmt.Print(strings.Replace(string(buf), pattern, rrep, -1))
		case proc == PROC_DEFAULTLEN:
			fmt.Print(string(buf) + NEWLINE)
		case proc == PROC_NOOP:
			fmt.Print(string(buf))
		case proc == PROC_REGEX:
			panic("TODO: implement")
		case proc == PROC_FUNC:
			panic("TODO: implement")
		}
		if err != nil && err != io.EOF {
			panic(err)
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		panic("Must provide pattern, provided empty string instead")
	}
	pattern := strings.Join(args, " ")
	pagedStdin(pattern, PROC_TEXT, BUFLEN)

}
