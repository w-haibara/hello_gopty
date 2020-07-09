package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/creack/pty"
)

func main() {
	ptmx, tty, err := pty.Open()
	if err != nil {
		panic(err)
	}
	defer ptmx.Close()
	defer tty.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")
		scanner.Scan()
		input := scanner.Bytes()

		ptmx.Write(input)

		buf := make([]byte, 1024)
		n, err := ptmx.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buf[:n]))

		if bytes.Equal(input, []byte("quit")) {
			fmt.Println("Bye")
			return
		}
	}
}
