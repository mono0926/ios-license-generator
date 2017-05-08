package reader

import (
	"io/ioutil"
	"strings"
)

func read(file string) (string, error) {
	bs, e := ioutil.ReadFile(file)
	if e != nil {
		return "", e
	}
	return string(bs), nil
}

func ReadLines(file string) ([]string, error) {
	s, e := read(file)
	if e != nil {
		return nil, e
	}
	return strings.Split(s, "\n"), nil
}
