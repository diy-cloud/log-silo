package main

import (
	"os"
	"strconv"

	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
)

func ReadParquetEnv(pw *writer.ParquetWriter) {
	pw.RowGroupSize = 256 * 1024 * 1024
	pw.PageSize = 4 * 1024
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	if value := os.Getenv("PARQUET_ROW_GROUP_SIZE"); value != "" {
		if v, err := strconv.Atoi(value); err == nil {
			pw.RowGroupSize = int64(v)
		}
	}
	if value := os.Getenv("PARQUET_PAGE_SIZE"); value != "" {
		if v, err := strconv.Atoi(value); err == nil {
			pw.PageSize = int64(v)
		}
	}
	if value := os.Getenv("PARQUET_COMPRESSION_TYPE"); value != "" {
		switch value {
		case "NONE":
			pw.CompressionType = parquet.CompressionCodec_UNCOMPRESSED
		case "SNAPPY":
			pw.CompressionType = parquet.CompressionCodec_SNAPPY
		case "GZIP":
			pw.CompressionType = parquet.CompressionCodec_GZIP
		case "LZO":
			pw.CompressionType = parquet.CompressionCodec_LZO
		case "BROTLI":
			pw.CompressionType = parquet.CompressionCodec_BROTLI
		case "ZSTD":
			pw.CompressionType = parquet.CompressionCodec_ZSTD
		}
	}
}
