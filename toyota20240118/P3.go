package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P3() {
	var q int
	fmt.Scan(&q)

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)

	now := 0
	ls := make([]int, 0)
	m := 0

	for i := 0; i < q; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		command, _ := strconv.Atoi(parts[0])

		switch command {
		case 1:
			l, _ := strconv.Atoi(parts[1])
			ls = append(ls, now)
			now += l
		case 2:
			m++
		case 3:
			d, _ := strconv.Atoi(parts[1])
			ans := (ls[d+m-1] - ls[m])
			fmt.Println(ans)
		}

	}

}
