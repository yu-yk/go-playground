package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the aVeryBigSum function below.
func aVeryBigSum(ar []int64) string {
	var digit []int8
	for _, n := range ar {
		// process each digit
		carry := int8(0)
		for i := 0; n > 0; i++ {
			if i >= len(digit) {
				digit = append(digit, 0)
			}
			d := n % 10
			s := int8(d) + digit[i] + carry
			if s > 10 {
				carry++
				digit[i] = s - 10
			} else {
				digit[i] = s
				carry = 0
			}
			n = n / 10
		}
		if carry > 0 {
			digit = append(digit, carry)
		}
	}
	for i := len(digit) - 1; i >= 0; i-- {
		fmt.Print(digit[i])
	}
	return ""
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout := os.Stdout

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int64

	for i := 0; i < int(arCount); i++ {
		arItem, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		ar = append(ar, arItem)
	}

	result := aVeryBigSum(ar)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
