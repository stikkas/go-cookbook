package main

import (
	"bytes"
	"fmt"
	"io"
	"stikkas.ru/go-cookbook/ch1/interfaces"
)

func main() {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}
	fmt.Println("out bytes buffer =", out.String())
	fmt.Print("stdout on PipeExample = ")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
	train()
}

func train() {
	r, w := io.Pipe()

	go func() {
		w.Write([]byte("Привет, доброе утро! Hello, how are you"))
		w.Close()
	}()
	data := make([]byte, 100)
	count, err := r.Read(data)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data[0:count]))
}
