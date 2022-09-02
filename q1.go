package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)


func incrementMap(ref map[int]int) {
	start := time.Now()
	for k, v := range ref {
		ref[k] = v + 1
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func incrementSliceOrArray(ref []int) {
	start := time.Now()
	for k, v := range ref {
		ref[k] = v + 1
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func main() {
	args := os.Args
	x, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	length := x + 1

	start := time.Now()
	nummap := make(map[int]int)
	for i := 0; i < length; i++ {
		nummap[i] = i
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	start = time.Now()
	numarray := make([]int, length)
	for i := 0; i < length; i++ {
		nummap[i] = i
	}
	elapsed = time.Since(start)
	fmt.Println(elapsed)

	start = time.Now()
	numslice := []int{}
	for i := 0; i < length; i++ {
		nummap[i] = i
	}
	elapsed = time.Since(start)
	fmt.Println(elapsed)
	
	
	incrementMap(nummap)

	incrementSliceOrArray(numarray)

	incrementSliceOrArray(numslice)

}
