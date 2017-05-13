package syscall

import (
	"os"
	"testing"
)

func TestPread(t *testing.T) {
	var err error
	if err = os.Mkdir("./.testing", 0755); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll("./.testing")

	var f *os.File
	if f, err = os.Create("./.testing/file.txt"); err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	if _, err = f.WriteString("hello world!"); err != nil {
		t.Fatal(err)
	}

	var (
		buf [32]byte
		n   int

		fd = int(f.Fd())
	)

	if n, err = Pread(fd, buf[:], 0); err != nil {
		t.Fatal(err)
	}

	if str := string(buf[:n]); str != "hello world!" {
		t.Fatalf("invalid response\nExpected: %s\nReturned: %s\n", "hello world!", str)
	}

	if n, err = Pread(fd, buf[:], 6); err != nil {
		t.Fatal(err)
	}

	if str := string(buf[:n]); str != "world!" {
		t.Fatalf("invalid response\nExpected: %s\nReturned: %s\n", "hello world!", str)
	}
}
