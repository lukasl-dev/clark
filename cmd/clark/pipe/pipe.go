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
