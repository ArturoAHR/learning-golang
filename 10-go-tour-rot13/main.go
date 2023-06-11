package main

// https://go.dev/tour/methods/23

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(byteArray []byte) (int, error) {
	bytes, err := reader.r.Read(byteArray)

	for i := 0; i < len(byteArray); i++ {

		if (byteArray[i] < 97 || byteArray[i] > 122) && (byteArray[i] < 65 || byteArray[i] > 90) {
			continue
		}

		var firstLetter byte

		if byteArray[i] >= 97 {
			firstLetter = 97
		} else {
			firstLetter = 65
		}

		byteArray[i] = firstLetter + ((byteArray[i] - firstLetter + 13) % 26)
	}

	if err == io.EOF {
		return 0, io.EOF
	}
	return bytes, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
