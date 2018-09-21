package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	n := nextLine()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = nextLine()
	}

	maxv := -100000000
	minv := arr[0]
	for j := 1; j < n; j++ {
		if maxv < arr[j] - minv {
			maxv = arr[j] - minv
		}
		if minv > arr[j] {
			minv = arr[j]
		}
	}
	fmt.Println(maxv)
}
