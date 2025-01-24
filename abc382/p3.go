package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	for i := 0; i < n; i++ {
		bishokudo[i], _ = strconv.Atoi(parts1[i])
		if i != 0 && bishokudo[i-1] < bishokudo[i] {
			bishokudo[i] = bishokudo[i-1]
		}
	}

	line2, _ := reader.ReadString('\n')
	parts2 := strings.Fields(strings.TrimSpace(line2))
	umami := make([]int, m)
	for i := 0; i < m; i++ {
		umami[i], _ = strconv.Atoi(parts2[i])
	}

	for i := 0; i < m; i++ {
		a := umami[i]

		i := sort.Search(n, func(i int) bool {
			return bishokudo[i] <= a
		})
		if i == n {
			fmt.Println(-1)
			continue
		}

		fmt.Println(i + 1)

	}

}
