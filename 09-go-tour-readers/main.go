package main

// https://go.dev/tour/methods/22

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(byteArray []byte) (int, error) {
	for i := 0; i < len(byteArray); i++ {
		byteArray[i] = 65
	}
	return len(byteArray), nil
}

func main() {
	reader.Validate(MyReader{})
}
