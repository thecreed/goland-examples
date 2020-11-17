package main

import (
	"html/template"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	checkError(err)
	err = tmpl.Execute(os.Stdout, data)
	checkError(err)

}

func main() {
	blabla := map[string]string{"gogo": "haha"}
	ahahah := []int{12, 43, 23}

	type jjjj struct {
		Name   string
		Age    int
		Weight float64
		Option struct {
			Issues map[string]string
		}
	}

	joe := jjjj{Name: "joe due", Age: 44, Weight: 182.3}
	joe.Option.Issues = map[string]string{"aa": "bb"}
	executeTemplate("showing slice values {{range .}} item - {{.}} {{end}}  \n", ahahah)
	executeTemplate("showing map values {{range $k,$val := .}} item - {{$k}} - {{$val}} {{end}} \n", blabla)
	executeTemplate("bobobo name: {{.Option.Issues}} \n", joe)

}
