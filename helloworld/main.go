package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a int
	fmt.Fscanln(reader, &a)

	inputNums, _ := reader.ReadString('\n')

	strSlice := strings.Split(strings.TrimSpace(inputNums), " ")
	intSlice := make([]int, 0, a)

	for _, s := range strSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			continue
		}
		intSlice = append(intSlice, num)
	}

	var b int
	fmt.Fscanln(reader, &b)

	count := 0

	for i := 0; i < a; i++ {
		if intSlice[i] == b {
			count += 1
		}
	}

	fmt.Fprint(writer, count)

	writer.Flush()
}
