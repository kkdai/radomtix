package main

import (
	"flag"
)

func main() {
	var numTickets int
	flag.IntVar(&numTickets, "n", 10, "Number of tickets (default: 10)")
	flag.Parse()

}
