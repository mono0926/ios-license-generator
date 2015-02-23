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

func GenerateLicense(t string, title string, body string) string {
    l := License{title, body}
    return generateImpl(t, l)
}

func generateImpl(t string, item interface{}) string {
    tpl := template.Must(template.ParseFiles(t))
    var b bytes.Buffer
    tpl.Execute(&b, item)
    s := b.String()
    return s

}