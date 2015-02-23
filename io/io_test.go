package reader
import (
    "testing"
    "fmt"
)
func TestParse(t *testing.T) {
    v, e := read("../template/License.plist")
    if e != nil {
        t.Errorf("error: %s", e)
    } else {
        fmt.Println(v)
    }
}