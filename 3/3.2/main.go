package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextLineArray() []int {
	sc.Scan()
	arr := strings.Split(sc.Text(), " ")
	ret := []int{}
	for _, v := range arr {
		i, _ := strconv.Atoi(v)
		ret = append(ret, i)
	}
	return ret
}

func output(arr []int) {
	str := ""
	for i := 0; i < len(arr); i++ {
		str += fmt.Sprintf("%d", arr[i])
		if i != len(arr) - 1 {
			str += " "
		}
	}
	fmt.Println(str)
}

func main() {
	n := nextLine()
	a := nextLineArray()

	output(a)
	for i := 1; i < n; i++ {
		v := a[i]
		j := i - 1
		for j >= 0 && a[j] > v {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = v
		output(a)
	}
}
