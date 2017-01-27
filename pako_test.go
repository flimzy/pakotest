package pako

import (
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func writerTest(w io.Writer, file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := io.Copy(w, f); err != nil {
		fmt.Printf("copy error: %s\n", err)
		panic(err)
	}

}

func BenchmarkStandardZlib1b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zw := zlib.NewWriter(ioutil.Discard)
		defer zw.Close()
		writerTest(zw, "testdata/1b.txt")
	}
}

func BenchmarkStandardZlib1kb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zw := zlib.NewWriter(ioutil.Discard)
		defer zw.Close()
		writerTest(zw, "testdata/1kb.txt")
	}
}

func BenchmarkStandardZlib1mb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zw := zlib.NewWriter(ioutil.Discard)
		defer zw.Close()
		writerTest(zw, "testdata/1mb.txt")
	}
}

func BenchmarkPakoZlib1b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zw := NewWriter(ioutil.Discard)
		defer zw.Close()
		writerTest(zw, "testdata/1b.txt")
	}
}

func BenchmarkPakoZlib1kb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zw := NewWriter(ioutil.Discard)
		defer zw.Close()
		writerTest(zw, "testdata/1kb.txt")
	}
}

func BenchmarkPakoZlib1mb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zw := NewWriter(ioutil.Discard)
		defer zw.Close()
		writerTest(zw, "testdata/1mb.txt")
	}
}
