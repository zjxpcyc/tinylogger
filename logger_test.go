package tinylogger_test

import (
	"os"
	"strings"
	"testing"

	"github.com/zjxpcyc/tinylogger"
)

func TestLogger(t *testing.T) {
	f, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.Fatal(err.Error())
	}

	var l tinylogger.LogService = tinylogger.NewLogger(f)

	l.Info("This is a success string")
	l.Error("This is a error string")

	f.Sync()
	f.Close()

	f2, err := os.Open("test.log")
	if err != nil {
		t.Fatal(err.Error())
	}

	var b []byte
	_, err = f2.Read(b)
	if err != nil {
		t.Fatalf("Can not read file: " + err.Error())
	}

	s := string(b)
	if strings.Index(s, "This is a success string") == 0 {
		t.Fatalf("Write Info message fail")
	}

	if strings.Index(s, "This is a error string") == 0 {
		t.Fatalf("Write Error message fail")
	}

	// l2 := tinylogger.NewLogger()
	// l2.Info("That is a success string")
	// l2.Error("That is a error string")

	// t.Fatalf("abc....")
}
