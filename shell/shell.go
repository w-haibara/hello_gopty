package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func main() {
	c := exec.Command("bash")
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}

	fmt.Println("-------------------------------------")

	go io.Copy(f, os.Stdin)
	io.Copy(os.Stdout, f)
}
