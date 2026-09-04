// Question 7.1 — Interfaces as contracts
//
// type ByteCounter int
// Implement io.Writer: Write appends to the count of bytes.
// Pass *ByteCounter to fmt.Fprintf and print the count.
//
// Extra: dump(w io.Writer, s string) and call it with os.Stdout and a bytes.Buffer.
//
// You will need: interface satisfaction, io.Writer, fmt.Fprintf.
//
// Run:
//   go run 7.1-bytecounter.go

package main

import "fmt"

type myBuffer struct {
	n int
	p []byte
}

func (b *myBuffer) Write(p []byte) (int, error) {
	b.n += len(p)
	b.p = append(b.p, p...)
	return len(p), nil
}

func (b *myBuffer) String() string {
	return string(b.p)
}

func (b *myBuffer) Size() int {
	return b.n
}

func main() {
	var b myBuffer
	fmt.Fprintf(&b, "hello")
	fmt.Println(b.String())
	fmt.Println(b.Size())

	fmt.Println("--------------------------------")

	fmt.Fprintf(&b, " world")
	fmt.Println(b.String())
	fmt.Println(b.Size())

}
