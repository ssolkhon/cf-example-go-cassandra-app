package models

type Record struct {
	Key   string
	Value string
}

func GetRecord(key string) *Record {
	return &Record{
		Key:   "Hello",
		Value: "World",
	}
}
