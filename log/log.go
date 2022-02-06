package log

type Log struct {
	UnixTime int64  `parquet:"name=unix_time, type=INT64" json:"unix_time"`
	AppID    int32  `parquet:"name=app_id, type=INT32" json:"app_id"`
	Level    int32  `parquet:"name=level, type=INT32" json:"level"`
	Message  string `parquet:"name=message, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"message"`
}
