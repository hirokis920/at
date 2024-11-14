package toyota202411

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type cell struct {
	x int64
	a int64
}

func main() {

	var n, m int64
	fmt.Scan(&n, &m)

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	line, _ := reader.ReadString('\n')
	parts1 := strings.Fields(line)
	line2, _ := reader.ReadString('\n')
	parts2 := strings.Fields(line2)

	cells := make([]cell, m)
	var sum int64 = 0
	for i := 0; i < int(m); i++ {
		x, _ := strconv.ParseInt(parts1[i], 10, 64)
		a, _ := strconv.ParseInt(parts2[i], 10, 64)
		cells[i] = cell{x: x, a: a}
		sum += a
	}

	sort.Slice(cells, func(i, j int) bool {
		return cells[i].x < cells[j].x
	})
	cells = append(cells, cell{n + 1, 1})

	var ans int64 = 0

	var px int64 = 0
	var num int64 = 1
	for _, cell := range cells {
		L := (cell.x - px)
		carry := num - L

		if carry < 0 {
			fmt.Println(-1)
			return
		}

		ans += (L - 1) * L / 2
		ans += L * carry
		num = carry + cell.a
		px = cell.x
	}

	if num != 1 {

	}

	fmt.Println(ans)
}
