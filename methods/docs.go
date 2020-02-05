package methods

import (
	"fmt"
	"ioutil"
)

/*
func gendocs() {
	var jsondata []Colls
} */
//Docreadcollect reads the Json data from path
func Docreadcollect() {
	data, err := ioutil.ReadFile("/Users/athul/Downloads/pwclc.json")
	if err != nil {
		fmt.Print(err)
	}
	var jsondata []Colls
	const tmpl = `The name is {{.Name}}.
{{range .Request}}
    His email id is {{.}}
{{end}}
`
}
