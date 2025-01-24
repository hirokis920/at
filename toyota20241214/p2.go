package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func p2() {

	var n, r int
	fmt.Scan(&n, &r)

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		parts1 := strings.Fields(strings.TrimSpace(line))
		if parts1[0] == "1" {
			if r >= 1600 && r <= 2799 {
				p, _ := strconv.Atoi(parts1[1])
				r += p
			}
		}
		if parts1[0] == "2" {
			if r >= 1200 && r <= 2399 {
				p, _ := strconv.Atoi(parts1[1])
				r += p
			}
		}

	}

	fmt.Println(r)

}
