package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/snowmerak/log-silo/log"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/writer"
)

func main() {
	if err := os.MkdirAll(filepath.Join(".", "logs"), 0777); err != nil {
		panic(err)
	}

	f, err := local.NewLocalFileWriter(filepath.Join(".", "log", time.Now().Format(time.RFC3339)+".parquet"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pw, err := writer.NewParquetWriter(f, new(log.Log), 8)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := pw.WriteStop(); err != nil {
			panic(err)
		}
	}()

	ReadParquetEnv(pw)

	RunReceiver(pw)
}
