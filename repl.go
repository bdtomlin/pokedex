package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func startRepl(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	for {
		printPrompt(w)
		scanner.Scan()
		execCommand(scanner.Text(), w)
	}
}

func printPrompt(w io.Writer) {
	fmt.Fprintln(w)
	fmt.Fprint(w, "pokedex > ")
}

func execCommand(cmd string, w io.Writer) {
	cmd = normalizeCmd(cmd)
	cmds := cliCommands()
	if _, ok := cmds[cmd]; !ok {
		fmt.Fprintln(w, "invalid command")
	}
	cmds[cmd].callback(w)
}

func normalizeCmd(cmd string) string {
	cmd = strings.ToLower(cmd)
	cmd = strings.Fields(cmd)[0]
	return cmd
}
