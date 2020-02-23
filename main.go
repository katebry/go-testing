package main

import (
	s "github.com/katebry/go-testing/Add"
	g "github.com/katebry/go-testing/GreeterFunction"
	r "github.com/katebry/go-testing/Iterate"
)

func main() {
	g.Greeter("kate", "German")
	s.Add(30, 39)
	r.Repeat("e")
}
