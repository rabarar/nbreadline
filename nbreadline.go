package nbreadline

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Reader struct {
	cmd      string
	err      chan error
	data     chan string
	ctrl     chan bool
	prompt   string
	sentinal byte
}

func (r *Reader) New() {
	r.err = make(chan error)
	r.ctrl = make(chan bool)
	r.data = make(chan string)
	r.sentinal = '\n'
	r.prompt = "> "

	go r.readLine()
}

func (r *Reader) Close() {
	// This will cause a deadlock - there is no way to close the routine
	// given that the readline blocks.
	//r.ctrl <- true
}

func (r *Reader) ReadLine() (string, error) {

	select {
	case cmd := <-r.data:
		return cmd, nil
	case err := <-r.err:
		return "", err
	default:
		return "", errors.New("Unknown state")
	}
}

func (r *Reader) readLine() {
	reader := bufio.NewReader(os.Stdin)

	for {
		select {
		case ctrl := <-r.ctrl:
			if ctrl {
				return
			}

		default:
			fmt.Printf(r.prompt)
			s, err := reader.ReadString(r.sentinal)
			if err != nil {
				r.err <- err
			} else {
				r.data <- s
			}
		}
	}
}
