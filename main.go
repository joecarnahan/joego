package main

import (
	"flag"
	"github.com/joecarnahan/joego/euler"
	"github.com/joecarnahan/joego/runner"
)

func main() {
	var reps = flag.Int("reps", 100, "Number of times to run the function")
	flag.Parse()
	runner.RunInt(euler.Problem001, euler.Problem001Default, *reps)
}
