package benchmarks

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var testString = "this is a long string of text that will be written to reader to test the performance of the readers and writers functions"

func readAllFile() {
	f, err := os.Open("twitchChat.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
}

func readFile() {
	f, err := os.Open("twitchChat.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	longBuf := make([]byte, 64)
	if _, err = io.ReadFull(f, longBuf); err != nil {
		fmt.Println("error:", err)
	}
}

func stringReader() {
	r := strings.NewReader(testString)
	buf := make([]byte, 64)
	if _, err := r.Read(buf); err != nil {
		fmt.Println("error:", err)
	}
}

func byteReader() {
	r := bytes.NewReader([]byte(testString))
	buf := make([]byte, 64)
	if _, err := r.Read(buf); err != nil {
		fmt.Println("error:", err)
	}
}

func stringBuilderWriteString() {
	var b strings.Builder
	b.WriteString(testString)
}

func ioCopy() {
	r := strings.NewReader(testString)
	b := new(bytes.Buffer)
	if _, err := io.Copy(b, r); err != nil {
		fmt.Println("error:", err)
	}
}

func byteWriter() {
	var b bytes.Buffer
	b.Write([]byte(testString))
}
