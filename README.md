# log-silo

A simple log store using parquet.

## environment variable

### PARQUET_ROW_GROUP_SIZE

default PARQUET_ROW_GROUP_SIZE = 268435456

row size is 256mb

### PARQUET_PAGE_SIZE

default PAEQUET_PAGE_SIZE = 8192

page size is 8kb

### PARQUET_COMPRESSION_TYPE

- NONE
- SNAPPY
- GZIP
- LZO
- BROTLI
- ZSTD

### SERVER_TYPE

- nats

#### NATS_URL

default NATS_URL = log

#### NATS_SUBJECT

- kafka

#### KAFKA_TOPIC

dafault KAFKA_TOPIC = log

#### KAFKA_PARTITION

default KAFKA_PARTITION = 0

#### KAFKA_NETWORK

default KAFKA_NETWORK = tcp

#### KAFKA_URL

defaul KAFKA_URL = localhost:9092

- http