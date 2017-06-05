// Ignore file as this is only used during 'go generate'
// +build ignore

// This program generates the webpage handlers
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"time"
)

var resources = map[string]string{
	"/": "rootpage.htm",
}

func main() {
	for res := range resources {
		fmt.Println("Generating ", resources[res])

		rawContent, err := ioutil.ReadFile(resources[res])
		if err != nil {
			panic(err)
		}

		f, err := os.Create(resources[res] + ".go")
		if err != nil {
			panic(err)
		}

		handlerTemplate.Execute(f, struct {
			Timestamp  time.Time
			Filename   string
			Content    string
			HandleFunc string
		}{
			Timestamp:  time.Now(),
			Filename:   resources[res],
			HandleFunc: strings.Replace(resources[res], ".", "_", -1),
			Content:    string(rawContent),
		})
	}
}

var handlerTemplate = template.Must(template.New("").Parse(`// DO NOT EDIT
// FILE WAS AUTOMATICALLY GENERATED ({{ .Timestamp }})
// EDIT '{{ .Filename }}' INSTEAD AND RERUN 'go generate'

package main

import (
  "net/http"
)

func {{ .HandleFunc }}(w http.ResponseWriter, req *http.Request) {
  w.Write(content)
}

var content = []byte(` + "`{{ .Content }}`)"))
