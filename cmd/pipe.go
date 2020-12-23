package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func readPipe(timeoutMS int64) []byte {
	var pipe []byte
	timeout := make(chan bool, 1)

	defer close(timeout)

	go func() {
		b, err := ioutil.ReadAll(os.Stdin)

		if err != nil {
			log.Fatalln(err.Error())
		}

		pipe = b
	}()

	go func() {
		time.Sleep(time.Duration(timeoutMS) * time.Millisecond)
		timeout <- true
	}()

	select {
	case <-timeout:
	}

	pipe = bytes.ReplaceAll(pipe, []byte("\n"), nil)
	pipe = bytes.ReplaceAll(pipe, []byte("\r"), nil)

	return pipe
}
