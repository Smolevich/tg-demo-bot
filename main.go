package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/smolevich/tg-demo-bot/storage"
)

func main() {
	storage, err := storage.NewPgxStorage(os.Getenv("DB_DSN"), 10, 10, time.Duration(0))
	if err != nil {
		log.Fatal(err)
	}
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

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
		b.Send(m.Chat, "Hello, "+m.Sender.Username)
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
		log.Printf("Chat id %d, name %s, sender %s, text %s", m.Chat.ID, m.Chat.Title, m.Sender.Username, m.Text)
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
	<-signalCh
	fmt.Println("Stopping bot...")
}
