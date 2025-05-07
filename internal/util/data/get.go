package data

import "os"

func Get(key string) (string, error) {
	data, error := os.ReadFile(datafile(key))
	return string(data), error
}
