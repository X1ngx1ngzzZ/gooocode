package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tmpl := `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
            <p>{{.}}</p>
        {{end}}
    <ul>
        {{range .Options}}
            <li> <a href="/{{.Arc}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>
</html>
`
	tpl = template.New("tmpl.html")
	tpl, _ = tpl.Parse(tmpl)
}

type Story map[string]Chapter

type Chapter struct {
	Title     string   `json:"title"`
	Paragraph []string `json:"story"`
	Options   []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

//返回一个路由函数
func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}
