package parser

import (
	"log"
	"testing"
)

func TestParse(t *testing.T) {
	if match, lib, license := Fetch("[SwiftyJSON]:https://github.com/SwiftyJSON/SwiftyJSON/blob/master/LICENSE"); match {
		log.Println(lib)
		log.Println(license)
	} else {
		t.Errorf("error")
	}

	if match, lib, license := Fetch("[OAuth2]:https://github.com/p2/OAuth2/blob/2.0.0/LICENSE.txt"); match && license != "Not Found" {
		log.Println(lib)
		log.Println(license)
	} else {
		t.Errorf("error")
	}

	if match, lib, license := Fetch("[SQLite.swift]:https://github.com/stephencelis/SQLite.swift/blob/master/LICENSE.txt"); match {
		log.Println(lib)
		log.Println(license)
	} else {
		t.Errorf("error")
	}

	if match, lib, license := Fetch("[ios-license-generator]:https://github.com/mono0926/ios-license-generator/blob/master/LICENSE"); match {
		log.Println(lib)
		log.Println(license)
	} else {
		t.Errorf("error")
	}

	if match, lib, license := Fetch("hoge"); match {
		t.Errorf("error: (lib: %s, license: %s)", lib, license)
	}
	if match, lib, license := Fetch("[SwiftyJSON]:hoge"); match {
		t.Errorf("error: (lib: %s, license: %s)", lib, license)
	}

}
