package main

import (
	"log"
	"os"
	"strconv"
	"math/rand"
	"sort"
)

func main() {
	args := os.Args
	x, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	length := x + 1

	arr := make([]int, length)
	var sli sort.IntSlice
	for i := 0; i < length; i++ {
		arr[i] = rand.Int()
		sli[i] = rand.Int()
	}

	sort.Sort(arr)
	sort.Sort(sli)
	sort.Stable(arr)
	sort.Stable(sli)





}
