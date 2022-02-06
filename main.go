package main

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/snowmerak/log-silo/log"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/writer"
)

func main() {
	path := os.Getenv("PATH")
	if path == "" {
		path = filepath.Join(".", "logs")
	}

	if err := os.MkdirAll(path, 0777); err != nil {
		panic(err)
	}

	f, err := local.NewLocalFileWriter(filepath.Join(path, time.Now().Format(time.RFC3339)+".parquet"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pw, err := writer.NewParquetWriter(f, new(log.Log), int64(runtime.NumCPU()))
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
