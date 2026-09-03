package fsserver

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed assets
var assets embed.FS

func ServeTemplate(w http.ResponseWriter, name string, data map[string]interface{}) {
	t, err := readTemplate(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bootstrapData, err := assets.ReadFile("assets/bootstrap.min.css")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newData := map[string]interface{}{}
	for k, v := range data {
		newData[k] = v
	}
	newData["bootstrapCSS"] = template.CSS(string(bootstrapData))

	t.Execute(w, newData)
}

func readTemplate(name string) (*template.Template, error) {
	data, err := assets.ReadFile("assets/" + name + ".html")
	if err != nil {
		return nil, err
	}
	t := template.New(name)
	return t.Parse(string(data))
}
