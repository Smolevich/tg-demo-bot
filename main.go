package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"

	"github.com/smolevich/tg-demo-bot/storage"
)

func main() {
	storage, err := storage.NewPgxStorage(os.Getenv("DB_DSN"), 10, 10, time.Duration(0))
	if err != nil {
		return
	}
	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}
	//defer conn.Close(context.Background())

	b, err := tb.NewBot(tb.Settings{
		URL:    "https://api.telegram.org",
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Start bot")

	b.Handle("/hello", func(m *tb.Message) {
		log.Println(m.Chat, m.Text, m.Sender.ID)
		b.Send(m.Sender, "Hello World!")
	})
	b.Handle(tb.OnText, func(m *tb.Message) {
		// all the text messages that weren't
		// captured by existing handlers
		err := storage.Exec(
			`INSERT INTO messages
			(chat_id, username, user_id, text)
			VALUES ($1, $2, $3, $4)`,
			m.Chat.ID,
			m.Sender.Username,
			m.Sender.ID,
			m.Text,
		)
		if err != nil {
			fmt.Println("Error", err)
		}
		log.Printf("Chat id %s, name %s, sender %s, text %s", m.Chat.ID, m.Chat.Title, m.Sender.Username, m.Text)
	})

	b.Handle(tb.OnPhoto, func(m *tb.Message) {
		// photos only
	})

	b.Handle(tb.OnChannelPost, func(m *tb.Message) {
		// channel posts only
	})

	b.Handle(tb.OnQuery, func(q *tb.Query) {
		// incoming inline queries
	})

	b.Start()
}
