/*
 *    Copyright 2021 lukasl-dev
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package pipe

import (
	"bytes"
	"io/ioutil"
	"os"
	"time"
	"unicode"
)

func Read(timeout time.Duration) ([]byte, error) {
	var dest []byte
	var err error

	deadline := make(chan bool, 1)

	go func() {
		dest, err = ioutil.ReadAll(os.Stdin)
		deadline <- true
	}()

	go func() {
		time.Sleep(timeout)
		deadline <- false
	}()

	defer close(deadline)

	<-deadline

	dest = bytes.TrimRightFunc(dest, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	return dest, nil
}

func Write(b []byte) error {
	_, err := os.Stdout.Write(b)
	return err
}

func WriteString(s string) error {
	_, err := os.Stdout.WriteString(s)
	return err
}
