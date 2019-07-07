package main

import (
	"flag"
	"fmt"
)

const (
	DefaultTickets int = 10
	DefaultLength  int = 6
)

func main() {
	var numTickets int
	var lenString int
	flag.IntVar(&numTickets, "n", DefaultTickets, "Number of tickets (default: 10)")
	flag.IntVar(&lenString, "l", DefaultLength, "Length of tickets invited code (default: 6)")
	flag.Parse()

	if numTickets < 0 {
		numTickets = DefaultTickets
	}

	if lenString < 0 {
		numTickets = DefaultLength
	}

	for numTickets > 0 {
		fmt.Println(RandStringBytes(lenString))
		numTickets--
	}
}
