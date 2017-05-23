package main

import (
	"log"

	"github.com/gomatic/lib_b"
	"github.com/gomatic/lib_c"
	"github.com/gomatic/lib_d"
)

func main() {
	m := lib_d.D{B: lib_b.B{}, C: lib_c.C{}}
	log.Printf("%#v", m)
}
