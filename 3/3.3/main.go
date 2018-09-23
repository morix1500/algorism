package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
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

	cnt := 0
	flg := true
	for flg {
		flg = false
		for i := n - 1; i > 0; i-- {
			if a[i] < a[i-1] {
				tmp := a[i]
				a[i] = a[i-1]
				a[i-1] = tmp
				flg = true
				cnt++
			}
		}
	}
	output(a)
	fmt.Println(cnt)
}
