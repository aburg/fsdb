/*
Copyright © 2025 Andreas Burg abu@andreasburg.net

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package util

import (
	"fmt"
	"os"
	"regexp"
)

func Set(key string, value string) error {
	keyfile := keyfile(key)

	return os.WriteFile(keyfile, []byte(value), 0644)
}

func Get(key string) (string, error) {
	keyfile := keyfile(key)

	data, err := os.ReadFile(keyfile)

	return string(data), err
}

func Unset(key string) error {
	keyfile := keyfile(key)
	return os.Remove(keyfile)
}

func keyfile(key string) string {
	match, err := regexp.MatchString("^\\w+(\\.\\w+)*$", key)
	fail(err)

	if !match {
		fmt.Println("That is not a valid key:", key)
		os.Exit(1)
	}

	return "." + key
}

func fail(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
