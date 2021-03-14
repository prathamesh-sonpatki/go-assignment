package main

import (
	"fmt"
	"os"
	"os/signal"
)

func cpuBound(n int) {
	f, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for {
		fmt.Fprintf(f, ".")
	}
}

func main() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	cpuBound(0)
	s := <-c
	fmt.Println("Got signal:", s)
}
