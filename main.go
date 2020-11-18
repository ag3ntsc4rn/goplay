package main

import (
	"io"
	"os"
)

func main(){
	f, err := os.Create("Hello.txt")
	if err != nil {
		panic((err))
	}
	_, err = io.WriteString(f, "Hello World!")
	if err != nil {
		panic((err))
	}

	_, err = io.WriteString(f, "More Text!")
	if err != nil {
		panic((err))
	}

	err = f.Close()
	if err != nil {
		panic((err))
	}

}