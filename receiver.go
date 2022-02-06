package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/segmentio/kafka-go"
	"github.com/snowmerak/log-silo/util/signal"
	"github.com/xitongsys/parquet-go/writer"
)

func RunReceiver(pw *writer.ParquetWriter) {
	serverType := os.Getenv("SERVER_TYPE")
	switch serverType {
	case "nats":
		RunNats(pw)
	case "kafka":
		RunKafka(pw)
	case "http":
		RunHTTP(pw)
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

	fmt.Println("connect on", natsURL, ", subject: ", subject)

	<-signal.NewTerminate()
}

func RunKafka(pw *writer.ParquetWriter) {
	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		topic = "log"
	}
	partition := 0
	if value := os.Getenv("KAFKA_PARTITION"); value != "" {
		v, err := strconv.Atoi(value)
		if err == nil {
			partition = v
		}
	}
	network := "tcp"
	if value := os.Getenv("KAFKA_NETWORK"); value != "" {
		network = value
	}
	url := "localhost:9092"
	if value := os.Getenv("KAFKA_URL"); value != "" {
		url = value
	}

	conn, err := kafka.DialLeader(context.Background(), network, url, topic, partition)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("connect on ", url, ", topic: ", topic, ", partition: ", partition)

	buf := make([]byte, 8*1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		l := new(Log)
		if err := json.Unmarshal(buf[:n], l); err != nil {
			l.AppID = -1
			l.Level = ERROR
			l.Message = err.Error()
			l.UnixTime = time.Now().Unix()
		}
		if err := pw.Write(l); err != nil {
			panic(err)
		}
	}
}

func RunHTTP(pw *writer.ParquetWriter) {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			l := new(Log)
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(l); err != nil {
				l.AppID = -1
				l.Level = ERROR
				l.Message = err.Error()
				l.UnixTime = time.Now().Unix()
			}
			if err := pw.Write(l); err != nil {
				panic(err)
			}
		}
	})

	url := "0.0.0.0:14219"
	if value := os.Getenv("HTTP_URL"); value != "" {
		url = value
	}

	fmt.Println("Listening on", url)

	if err := http.ListenAndServe(url, nil); err != nil {
		panic(err)
	}
}
