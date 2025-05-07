package data

import (
	"os"
)

func Set(key string, value string) error {
	return os.WriteFile(datafile(key), []byte(value), 0644)
}
