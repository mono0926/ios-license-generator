package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func Fetch(line string) (match bool, key string, license string) {
	if match, k, url := Parse(line); match {
		l := get(url)
		if l == "" {
			return false, k, ""
		}
		l = strings.Replace(l, "<", "", -1)
		l = strings.Replace(l, ">", "", -1)
		return true, k, l
	}
	return false, "", ""
}

func Parse(line string) (match bool, key string, url string) {
	log.Println(line)

	r, _ := regexp.Compile(`\[[\-\.\w]+\]:`)

	lib := r.FindString(line)
	fmt.Println(lib)
	if len(lib) > 0 {
		u := strings.Replace(line, lib, "", 1)
		lib = strings.Replace(lib, "[", "", 1)
		lib = strings.Replace(lib, "]:", "", 1)
		log.Printf("%s: %s", lib, u)

		u = strings.Replace(u, "https://github.com/", "https://raw.githubusercontent.com/", 1)
		u = strings.Replace(u, "/blob/", "/", 1)

		return true, lib, u
	}
	return false, "", ""
}

func get(url string) string {
	log.Println("url: " + url)
	r, e := http.Get(url)
	if e != nil {
		return ""
	}
	defer r.Body.Close()

	bs, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return ""
	}
	return string(bs)
}
