package aio

import (
	"fmt"
	//	"os"
	"bytes"
	"io"
	"os"
	"testing"
)

var ta = New()
var testBuf = bytes.NewBuffer(nil)

const (
	testFilePath = "./testing/declarationOfIndependence.txt"
)

func TestBasic(t *testing.T) {
	var (
		f     *File
		wf    *File
		oresp *OpenResp
	)

	aio := New()

	if oresp = <-aio.Open("./testing/helloWorld.txt"); oresp.Err != nil {
		t.Fatal(oresp.Err)
	}

	f = oresp.F

	if oresp = <-aio.OpenFile("./testing/testWrite.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600); oresp.Err != nil {
		t.Fatal(oresp.Err)
	}

	wf = oresp.F

	var (
		buf [32]byte
		n   int
		err error
	)

	if n, err = f.Read(buf[:]); err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(buf[:n]))

	if _, err = wf.Write([]byte("hai hai hai!")); err != nil {
		t.Fatal(err)
	}

	if err = wf.Close(); err != nil {
		t.Fatal(err)
	}

	if err = <-aio.Delete("./testing/testWrite.txt"); err != nil {
		t.Fatal(err)
	}

	or := <-ta.Open(testFilePath)
	if or.Err != nil {
		t.Fatal(err)
	}

	if _, err := io.Copy(testBuf, or.F); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkAIO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testBuf.Reset()
		or := <-ta.Open(testFilePath)
		if or.Err != nil {
			b.Fatal(or.Err)
		}

		if _, err := io.Copy(testBuf, or.F); err != nil {
			b.Fatal(err)
		}

		if err := or.F.Close(); err != nil {
			b.Fatal(err)
		}
	}

	b.ReportAllocs()
}

func BenchmarkStdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testBuf.Reset()
		f, err := os.Open(testFilePath)
		if err != nil {
			b.Fatal(err)
		}

		if _, err := io.Copy(testBuf, f); err != nil {
			b.Fatal(err)
		}

		if err := f.Close(); err != nil {
			b.Fatal(err)
		}
	}

	b.ReportAllocs()
}

func BenchmarkAIOPara(b *testing.B) {
	b.SetParallelism(64)
	b.RunParallel(func(pb *testing.PB) {
		buf := bytes.NewBuffer(nil)
		for pb.Next() {
			or := <-ta.Open(testFilePath)
			if or.Err != nil {
				b.Fatal(or.Err)
			}

			if _, err := io.Copy(buf, or.F); err != nil {
				b.Fatal(err)
			}

			if err := or.F.Close(); err != nil {
				b.Fatal(err)
			}

			buf.Reset()
		}
	})

	b.ReportAllocs()
}

func BenchmarkStdlibPara(b *testing.B) {
	b.SetParallelism(64)
	b.RunParallel(func(pb *testing.PB) {
		buf := bytes.NewBuffer(nil)
		for pb.Next() {
			f, err := os.Open(testFilePath)
			if err != nil {
				b.Fatal(err)
			}

			if _, err := io.Copy(buf, f); err != nil {
				b.Fatal(err)
			}

			if err := f.Close(); err != nil {
				b.Fatal(err)
			}

			buf.Reset()
		}
	})

	b.ReportAllocs()
}
