package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var i int = -1
	var file string

	//Create a collection of files to compress
	for i, file = range os.Args[1:] {
		wg.Add(1)

		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
		wg.Wait()
	}
	fmt.Printf("Compressed %d files\n", i+1)

}

func compress(filename string) error {
	//Open the file to be compressed
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	//Open the destination file with the .gz extension
	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	//gzip.NewWriter compresses the file and writes it to the destination file
	gzout := gzip.NewWriter(out)

	//io.Copy copies the contents of the file to the gzip writer
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}
