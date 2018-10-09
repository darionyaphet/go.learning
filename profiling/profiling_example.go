package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

// https://blog.golang.org/profiling-go-programs

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func startMemProfile() {
	if *memProfile != "" && *memProfileRate > 0 {
		runtime.MemProfileRate = *memProfileRate
	}
}
func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
}
