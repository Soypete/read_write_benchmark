package benchmarks

import (
	"bufio"
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var testString = "this is a long string of text that will be written to reader to test the performance of the readers and writers functions"

// Comapre performace of reading a complete stream vs reading a stream in chunks
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

func readFileBuf() {
	file, err := os.Open("twitchChat.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func readFromDB() {
	fileName := "twitchChat.db"
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM chat")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var message string
		rows.Scan(&id, &message)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// func readFile() {
// 	f, err := os.Open("twitchChat.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
// 	longBuf := make([]byte, 64)
// 	if _, err = io.ReadFull(f, longBuf); err != nil {
// 		fmt.Println("error:", err)
// 	}
// }

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

// func messageHandler(w http.ResponseWriter, r *http.Request) {
// 	// Set the content type header to text/plain
// 	w.Header().Set("Content-Type", "text/plain")

// 	_, err := fmt.Fprintf(w, testString)
// 	if err != nil {
// 		log.Println("Error writing message:", err)
// 	}
// }
