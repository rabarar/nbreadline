package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rabarar/nbreadline"
)

func reader() {

	var r nbreadline.Reader

	r.New()
	defer r.Close()

	for {

		cmd, err := r.ReadLine()
		if err == nil {
			fmt.Printf("received cmd: %s", cmd)
		} else {
			if err == io.EOF {
				fmt.Printf("received EOF\n")
				break
			}
		}
		time.Sleep(time.Millisecond * 50)

	}

}

func main() {
	reader()
	fmt.Printf("exiting...\n")
	time.Sleep(time.Millisecond * 2000)
	os.Exit(1)
}
