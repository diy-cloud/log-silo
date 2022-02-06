# log-silo

A simple log store using parquet.

## environment variable

### PATH

PATH is using for storing log data.

default PATH = ./logs

### PARQUET_ROW_GROUP_SIZE

default PARQUET_ROW_GROUP_SIZE = 268435456

row size is 256mb

### PARQUET_PAGE_SIZE

default PAEQUET_PAGE_SIZE = 8192

page size is 8kb

### PARQUET_COMPRESSION_TYPE

This is meaning how to compress parquet file.

- NONE
- SNAPPY
- GZIP
- LZO
- BROTLI
- ZSTD

### SERVER_TYPE

connecting nats server

- nats

#### NATS_URL

default NATS_URL = nats://127.0.0.1:4222

#### NATS_SUBJECT

default NATS_SUBJECT = log

---

- kafka

#### KAFKA_TOPIC

dafault KAFKA_TOPIC = log

#### KAFKA_PARTITION

default KAFKA_PARTITION = 0

#### KAFKA_NETWORK

default KAFKA_NETWORK = tcp

#### KAFKA_URL

defaul KAFKA_URL = localhost:9092

---

- http

#### HTTP_URL

default HTTP_URL = 0.0.0.0:14219
