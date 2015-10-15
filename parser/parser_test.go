package parser
import (
    "testing"
    "log"
)

func TestParse(t *testing.T) {
    if match, lib, url := Fetch("[SwiftyJSON]:https://github.com/SwiftyJSON/SwiftyJSON/blob/master/LICENSE"); match {
        log.Println(lib)
        log.Println(url)
    } else {
        t.Errorf("error")
    }

    if match, lib, url := Fetch("[OAuth2]:https://github.com/p2/OAuth2/blob/master/LICENSE.txt"); match {
        log.Println(lib)
        log.Println(url)
    } else {
        t.Errorf("error")
    }

    if match, lib, url := Fetch("[SQLite.swift]:https://github.com/stephencelis/SQLite.swift/blob/master/LICENSE.txt"); match {
        log.Println(lib)
        log.Println(url)
    } else {
        t.Errorf("error")
    }

    if match, lib, url := Fetch("[ios-license-generator]:https://github.com/mono0926/ios-license-generator/blob/master/LICENSE"); match {
        log.Println(lib)
        log.Println(url)
    } else {
        t.Errorf("error")
    }


    if match, lib, url := Fetch("hoge"); match {
        t.Errorf("error: (lib: %s, url: %s)", lib, url)
    }
    if match, lib, url := Fetch("[SwiftyJSON]:hoge"); match {
        t.Errorf("error: (lib: %s, url: %s)", lib, url)
    }

}