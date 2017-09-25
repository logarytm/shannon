package main

import (
	"fmt"
	"os"
	"io"
	"math"
)

type Histogram []int

func Entropy(total int, h Histogram) float64 {
	entropy := 0.0
	for _, occurences := range h {
		probability := float64(occurences) / float64(total)
		entropy += probability * math.Log2(probability)
	}
	// absolute value instead of additive inverse to avoid returning negative zero
	return math.Abs(entropy)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args[1:]

	var f *os.File
	const blocksize = 1024 * 2048

	h := make([]int, 256)
	if len(args) == 0 {
		f = os.Stdin
	} else {
		var err error
		f, err = os.Open(args[0])
		check(err)
		defer f.Close()
	}

	buf := make([]byte, blocksize)
	total := 0
	for {
		read, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		check(err)

		for _, b := range(buf[0:read]) {
			h[b] += 1
		}
		total += read
	}

	e := Entropy(total, h)
	fmt.Printf("Entropy: %.4f (%.2f%% of maximum entropy)\n", e, 100 * e / 8)
}
