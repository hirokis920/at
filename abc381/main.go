package abc381

import (
	"bufio"
	"fmt"
	"os"
)

func p1() {

	var n int64
	fmt.Scan(&n)

	if (n % 2) == 0 {
		fmt.Println("No")
		return
	}

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	line, _ := reader.ReadString('\n')
	parts1 := []rune(line)

	harf := int(n-1) / 2

	for i := 0; i < harf; i++ {
		if parts1[i] != '1' {
			fmt.Println("No")
			return
		}
	}

	if parts1[harf] != '/' {
		fmt.Println("No")
		return
	}
	for i := 0; i < harf; i++ {
		if parts1[harf+i+1] != '2' {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")

}
