package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const NEWLINE = "\n"

type Proc int

const (
	PROC_TEXT Proc = iota
	PROC_DEFAULTLEN
	PROC_NOOP
	PROC_REGEX
	PROC_FUNC
)

func addNewlines(pattern string, chunk string) string {
	return strings.Replace(chunk, pattern, pattern+NEWLINE, -1)
}

func pagedStdin(pattern string, proc Proc) {
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*4096)

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
			fmt.Printf(addNewlines(pattern, string(buf)))
		case proc == PROC_DEFAULTLEN:
			fmt.Printf(string(buf) + NEWLINE)
		case proc == PROC_NOOP:
			fmt.Printf(string(buf))
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
	pattern := flag.String("p", "", "will insert a new line after provided text expression")
	flag.Parse()
	fn := PROC_TEXT

	if *pattern == "" {
		panic("Must provide pattern, provided empty string instead")
	}

	pagedStdin(*pattern, fn)

}
