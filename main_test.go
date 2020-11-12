package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMeasure(t *testing.T) {
	args := []string{
		"--matrix", "1", "--matrix-size", "1000", "-t", "5s",
	}

	command := "stress-ng"
	output := bytes.NewBufferString("")

	if err := Measure(command, args, 1000 * time.Millisecond, output) ; err != nil {
		t.Errorf("Measure() : got %v", err)
		t.Error(err)
	}


	bytes, err := ioutil.ReadAll(output)
	if err != nil {
		t.Errorf("ioutil.Readall: got %v", err)
	}

	content := string(bytes)
	if want := "##FINISHED" ; !strings.Contains(content, want) {
		t.Errorf("invalid output, want : %q got %v", want, content)
	}

	if want := "Interval: 1\n" ; !strings.Contains(content, want) {
		t.Errorf("invalid output, want : %q got %v", want, content)
	}
}

func TestMeasureNilArgs(t *testing.T) {
	var args []string

	command := "ls"
	output := bytes.NewBufferString("")

	if err := Measure(command, args, 1000 * time.Millisecond, output) ; err != nil {
		t.Errorf("Measure() : got %v", err)
		t.Error(err)
	}


	bytes, err := ioutil.ReadAll(output)
	if err != nil {
		t.Errorf("ioutil.Readall: got %v", err)
	}

	content := string(bytes)
	if want := "##FINISHED" ; !strings.Contains(content, want) {
		t.Errorf("invalid output, want : %q got %v", want, content)
	}

	if want := "Interval: 1\n" ; !strings.Contains(content, want) {
		t.Errorf("invalid output, want : %q got %v", want, content)
	}
}

func TestMeasureDuration(t *testing.T) {
	var args []string

	command := "ls"
	output := bytes.NewBufferString("")

	if err := Measure(command, args, time.Duration(2) * time.Second, output) ; err != nil {
		t.Errorf("Measure() : got %v", err)
		t.Error(err)
	}


	bytes, err := ioutil.ReadAll(output)
	if err != nil {
		t.Errorf("ioutil.Readall: got %v", err)
	}

	content := string(bytes)
	if want := "##FINISHED" ; !strings.Contains(content, want) {
		t.Errorf("invalid output, want : %q got %v", want, content)
	}

	if want := "Interval: 2\n" ; !strings.Contains(content, want) {
		t.Errorf("invalid output, want : %q got %v", want, content)
	}
}

func Test_Main(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	file, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Errorf("some problem creating tmp file : %v",  err)

	}
	os.Args = []string{"procmem", file.Name(), "1", "sleep", "2"}

	main()

}