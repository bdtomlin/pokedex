package main

import (
	"bufio"
	"fmt"
	"io"
)

func startRepl(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	for {
		fmt.Fprintln(w)
		fmt.Fprint(w, "pokedex > ")
		scanner.Scan()
		cmd := scanner.Text()
		if err := execCommand(cmd, w); err != nil {
			fmt.Fprintln(w, err)
		}
	}
}
