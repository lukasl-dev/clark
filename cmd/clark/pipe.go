/*
 * Copyright 2021 lukasl-dev
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
