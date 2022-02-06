package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/snowmerak/log-silo/util/signal"
	"github.com/xitongsys/parquet-go/writer"
)

func RunReceiver(pw *writer.ParquetWriter) {
	serverType := os.Getenv("SERVER_TYPE")
	switch serverType {
	case "nats":
		RunNats(pw)
	case "kafka":
	case "http":
	}
}

func RunNats(pw *writer.ParquetWriter) {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	nc, _ := nats.Connect(natsURL)
	defer nc.Close()

	subject := os.Getenv("NATS_SUBJECT")
	if subject == "" {
		subject = "log"
	}

	sub, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		l := new(Log)
		if err := json.Unmarshal(msg.Data, l); err != nil {
			l.AppID = -1
			l.Level = ERROR
			l.Message = err.Error()
			l.UnixTime = time.Now().Unix()
		}
		if err := pw.Write(l); err != nil {
			panic(err)
		}
	})
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	<-signal.NewTerminate()
}
