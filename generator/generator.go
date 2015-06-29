package generator
import (
    "text/template"
    "bytes"
)

type Item struct {
     Name string
 }

type ItemList struct {
    Item string
}

type License struct {
    Title string
    Body string
}

func generateLicenseListItem(t string, name string) string {
    i := Item{name}
    return generateImpl(t, i)
}

func GenerateLicenseList(listPath string, itemPath string, names []string) string {
    r := ""
    for _, v := range names {
        r += generateLicenseListItem(itemPath, v)
    }
    i := ItemList{r}
    return generateImpl(listPath, i)
}

func GenerateLicense(src string, title string, body string) string {
    l := License{title, body}
    return generateImpl(src, l)
}

func generateImpl(src string, item interface{}) string {
    tpl := template.New("template")
    tpl.Parse(src)
    var b bytes.Buffer
    tpl.Execute(&b, item)
    s := b.String()
    return s

}