package data

import "os"

func Remove(key string) error {
	return os.Remove(datafile(key))
}
