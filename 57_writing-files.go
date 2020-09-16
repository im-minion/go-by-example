package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("test", d1, 0644)
	check(err)
	//For more granular writes, open a file for writing.

	f, err := os.Create("test2")
	check(err)
	//It’s idiomatic to defer a Close immediately after opening a file.

	defer f.Close()
	//You can Write byte slices as you’d expect.

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)
	//A WriteString is also available.

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	//Issue a Sync to flush writes to stable storage.

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
