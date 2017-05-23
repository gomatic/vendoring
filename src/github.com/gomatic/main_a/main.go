package main

import (
	"log"

	"github.com/gomatic/lib_b"
	"github.com/gomatic/lib_d"
)

func main() {
	m := lib_b.B{D: lib_d.D{}}
	log.Printf("%#v", m)
}
