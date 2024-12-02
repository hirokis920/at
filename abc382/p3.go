package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	line, _ := reader.ReadString('\n')
	parts1 := strings.Fields(strings.TrimSpace(line))
	bishokudo := make([]int, n)
	bishokumin := 200000
	for i := 0; i < n; i++ {
		bishokudo[i], _ = strconv.Atoi(parts1[i])
		if bishokumin > bishokudo[i] {
			bishokumin = bishokudo[i]
		}
	}
	line2, _ := reader.ReadString('\n')
	parts2 := strings.Fields(strings.TrimSpace(line2))
	umami := make([]int, m)

	for i := 0; i < m; i++ {
		umami[i], _ = strconv.Atoi(parts2[i])
	}

	eated := false
	for i := 0; i < m; i++ {
		eated = false
		a := umami[i]
		if a < bishokumin {
			fmt.Println(-1)
			continue
		}
		for j := 0; j < n; j++ {
			b := bishokudo[j]
			if a >= b {
				fmt.Println(j + 1)
				eated = true
				break
			}
		}
		if !eated {
			fmt.Println(-1)
		}
	}

}
