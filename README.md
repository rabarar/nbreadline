# nbreadline
Non-blocking readline() package for Golang

This is a simple package to implement a non-blocking readline in Go using channels and Go-routines

Example use:

	
	import "github.com/rabarar/nbreadline"


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


