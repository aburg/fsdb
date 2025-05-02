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

func Spawn(kind string, name string) (string, error) {
	assertValidSpawnName(name)

	err := os.Mkdir(name, 0755)
	if err != nil {
		fmt.Println("Could not create spawn folder")
		os.Exit(1)
	}
	os.Chdir(name)
	uuid, err := Identify(kind)
	os.Chdir("..")
	return uuid, err
}

func assertValidSpawnName(name string) {
	matched, err := regexp.MatchString("^[^\\.]\\w+$", name)
	if err != nil {
		fmt.Println("Could not check for valid spawn name")
		os.Exit(1)
	}
	if !matched {
		fmt.Println("Invalid name:", name)
		os.Exit(1)
	}
}
