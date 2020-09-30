package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	word := scanner.Text()

	scanner.Scan()
	input := scanner.Text()
	wordList := strings.Split(input, " ")

	m := make(map[rune]bool)
	for _, r := range word {
		m[r] = true
	}

	counter := 0

	for _, w := range wordList {
		wmap := make(map[rune]bool)
		for _, r := range w {
			wmap[r] = true
		}
		if len(wmap) != len(m) {
			continue
		}

		eq := reflect.DeepEqual(m, wmap)
		if eq {
			counter++
		}
	}

	fmt.Fprintln(os.Stdout, counter)
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	// set the split function by lines, default //會影響 Scan function 一次 scan 幾多野
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() { // 需要注意的是，Scan会将分割符号\n去除，如果Fprintf输出的话，不添加\n打印，会出现没有换行的现象
		fmt.Println(scanner.Text())
	}

	// Scanner在初始化的时候有设置一个maxTokenSize，
	// 这个值默认是MaxScanTokenSize = 64 * 1024，
	// 当一行的长度大于64*1024即65536之后，就会出现ErrTooLong这个错误
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func reader() {
	// reader := bufio.NewReader(os.Stdin)

	reader := bufio.NewReader(strings.NewReader("a\nb\nc\nd"))

	// 返回的结果是包含界定符本身的，上例中，输出结果有一空行就是'\n'本身(line携带一个'\n',printf又追加了一个'\n')
	line, _ := reader.ReadString('\n')
	fmt.Println(line) // two '\n'

	// 但是坑在于一定要做好分隔符的标示，如果你的文件文件中写的是"a\nb\nc\nd"
	// 使用ReadString的时候会漏掉d的输出，但是在最后一个line里面还有有数据的，所以需要自己手动的判断一下

}

func string() {
	str := "A B C D E"
	s := strings.Split(str, " ")
	strings.Join(s, " ")

	var buffer bytes.Buffer
	for i := 0; i < 1000; i++ {
		buffer.WriteString("a")
	}
	fmt.Println(buffer.String())

	var builder strings.Builder
	for i := 0; i < 1000; i++ {
		builder.WriteString("a")
	}
	fmt.Println(builder.String())

}
