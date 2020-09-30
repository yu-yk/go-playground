package main

import "fmt"

func main() {
	fmt.Println(addBinary("0", "0"))
}
func addBinary(a string, b string) string {
	var result string
	carry := 0
	p1 := len(a) - 1
	p2 := len(b) - 1

	for p1 >= 0 || p2 >= 0 {
		i1, i2, sum := 0, 0, 0
		if p1 >= 0 {
			i1 = int(a[p1] - '0')
		}
		if p2 >= 0 {
			i2 = int(b[p2] - '0')
		}
		sum += i1 + i2 + carry
		switch sum {
		case 0:
			result = "0" + result
			carry = 0
		case 1:
			result = "1" + result
			carry = 0
		case 2:
			result = "0" + result
			carry = 1
		case 3:
			result = "1" + result
			carry = 1
		}
		p1--
		p2--
	}

	if carry == 1 {
		result = "1" + result
	}

	return result
}
