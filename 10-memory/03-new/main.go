package main

import "fmt"

var allocatedBuffers = 0

func AllocateBuffer() *string {
	if allocatedBuffers < 3 {
		allocatedBuffers++
		return new(string)
	}
	return nil
}

func main() {
	var buffers []*string

	for {
		b := AllocateBuffer()
		if b == nil {
			break
		}

		buffers = append(buffers, b)
	}

	fmt.Println("Allocated", len(buffers), "buffers")
}
