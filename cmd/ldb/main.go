package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("ldb> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		out, err := processInput(text)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(*out)
	}
}

// TODO process input
func processInput(i string) (*string, error) {
	o := "ok"
	return &o, nil
}
