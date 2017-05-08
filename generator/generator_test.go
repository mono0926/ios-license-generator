package generator

import (
	"log"
	"testing"
)

func TestParse(t *testing.T) {
	s := generateLicenseListItem("../template/LicenseListItem.plist", "test1")
	log.Println(s)

	s = GenerateLicenseList("../template/LicenseList.plist", "../template/LicenseListItem.plist", []string{"a", "b"})
	log.Println(s)

	s = GenerateLicense("../template/License.plist", "hoge", "fuga")
	log.Println(s)

}
