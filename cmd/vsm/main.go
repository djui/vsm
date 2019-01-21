package main

import (
	"flag"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("")
	flag.Parse()
}
