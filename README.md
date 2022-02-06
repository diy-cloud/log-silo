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

- nats

connecting nats server

#### NATS_URL

default NATS_URL = nats://127.0.0.1:4222

#### NATS_SUBJECT

default NATS_SUBJECT = log

---

- kafka

connecting kafka server

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

listening http server

#### HTTP_URL

default HTTP_URL = 0.0.0.0:14219

send post request to "/"

---

## log data

`log` struct from `github.com/snowmerak/log-silo/log`.

```go
package log

type Log struct {
	UnixTime int64  `parquet:"name=unix_time, type=INT64" json:"unix_time"`
	AppID    int32  `parquet:"name=app_id, type=INT32" json:"app_id"`
	Level    int32  `parquet:"name=level, type=INT32" json:"level"`
	Message  string `parquet:"name=message, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"message"`
}
```

`log` has `unix_time`, `app_id`, `level` and `message`.

can represent to json.

```json
{
  "unix_time": 1598000,
  "app_id": 1,
  "level": 1,
  "message": "a log message"
}
```

`app_id` and `level` are customizable by your own team.

`message` is UTF8 string.

`unix_time` is unix time second. can use unix time nano.
