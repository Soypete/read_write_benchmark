package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fileName := "twitchChat.db"
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Creating table chat")
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS chat (id INTEGER PRIMARY KEY, user TEXT, message TEXT, timestamp TIMESTAMP)")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserting data into table chat")
	msg := "i'm sorry, i don't think i'm qualified to give advice on how to use a bot."
	_, err = db.Exec("INSERT INTO chat (user, message, timestamp) VALUES ('soy_un_bot',?,?)", msg, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	msg = "I'm going to eventually get myself one of these LLMS when it can more efficiently run on my i7 4790k and 970."
	_, err = db.Exec("INSERT INTO chat (user, message, timestamp) VALUES ('Cowtell', ?, ?)", msg, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	msg = "I'm a bit more hardware constrained in this new era of AI requirements."
	_, err = db.Exec("INSERT INTO chat (user, message, timestamp) VALUES ('Cowtell', ?, ?)", msg, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	msg = "I'm a bit more hardware constrained in this new era of AI requirements."
	_, err = db.Exec("INSERT INTO chat (user, message, timestamp) VALUES ('soy_llm_bot', ?, ?)", msg, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	msg = "I'd like to have at least 30 words per second generation for local use of my LLM."
	_, err = db.Exec("INSERT INTO chat (user, message, timestamp) VALUES ('Cowtell', ?, ?)", msg, time.Now())
	if err != nil {
		log.Fatal(err)
	}
}
