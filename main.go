package main

import (
	day1 "adventofcode/1"
	day2 "adventofcode/2"
	day3 "adventofcode/3"
	day4 "adventofcode/4"
	day5 "adventofcode/5"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func runner(f func()) {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	name = strings.TrimSuffix(name, ".Run")
	name = strings.TrimPrefix(name, "adventofcode/")
	fmt.Printf("############### --- Day %s --- ################\n", name)
	start := time.Now()
	f()
	fmt.Printf("Execution took: %dus\n", time.Since(start).Microseconds())
}

func main() {
	all := false
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "all" {
		all = true
	}

	start := time.Now()
	if all || args[0] == "day1" {
		runner(day1.Run)
	}
	if all || args[0] == "day2" {
		runner(day2.Run)
	}
	if all || args[0] == "day3" {
		runner(day3.Run)
	}
	if all || args[0] == "day4" {
		runner(day4.Run)
	}
	if all || args[0] == "day5" {
		runner(day5.Run)
	}

	fmt.Printf("\nTotal execution took: %dms\n", time.Since(start).Milliseconds())
}
