package main

import "os"

func main() {
	pd := newPokedex(os.Stdin, os.Stdout)
	startRepl(&pd)
}
