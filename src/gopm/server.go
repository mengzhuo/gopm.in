package gopm

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var repoTmpl *template.Template

func httpError(w http.ResponseWriter, code int, content string) {
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("URL:", r.URL.Path)

	if _, ok := r.URL.Query()["go-get"]; !ok {
		APIHandler(w, r)
		return
	}

	switch strings.Count(r.URL.Path, "/") {
	case 1: //  /guess.v1
		if strings.Contains(r.URL.Path, ".") {
			GuessHandler(w, r)
		}
	case 2: //  /gh/guess.v1
		SiteHandler(w, r)
	case 3: //  /gh/mengzhuo/guess.v1
		fallthrough
	default:
		SiteUserHandler(w, r)
	}
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func SiteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func githubHandler(w http.ResponseWriter, r *http.Request, urls []string) {
	fmt.Fprintf(w, "hello world")
}

func GuessHandler(w http.ResponseWriter, r *http.Request) {
}

func SiteUserHandler(w http.ResponseWriter, r *http.Request) {

	urls := strings.SplitN(r.URL.Path[1:], "/", 3)

	if len(urls) != 3 {
		w.Write([]byte("No valid url"))
		w.WriteHeader(403)
		log.Print("urls not enough:", r.URL.Path)
		return
	}
	var meta *MetaImport
	switch urls[0] {
	case "gh":
		fallthrough
	case "github":
		meta = &MetaImport{
			"gopm.in/" + "github/" + strings.Join(urls[1:], "/"),
			"git",
			"https://github.com/mengzhuo/bla.git",
		}
	case "gc":
		fallthrough
	case "google":
		fallthrough
	case "googlecode":

	case "bitbucket":

	default:
		log.Printf("No support site %#v", urls)
		w.Write([]byte("Not supported site"))
		w.WriteHeader(404)
		return
	}

	data := struct {
		Meta *MetaImport
	}{
		meta,
	}

	repoTmpl.Execute(w, data)
}

func init() {
	var err error
	repoTmpl, err = template.New("repo").Parse(GoImportTmpl)
	if err != nil {
		log.Fatal(err)
	}
}

type MetaImport struct {
	Prefix   string
	VCS      string
	RepoRoot string
}

const GoImportTmpl = `<html>
<head>
<meta content="{{.Meta.Prefix}} {{ .Meta.VCS }} {{ .Meta.RepoRoot }}" name="go-import" >
</head>
<body>
<ul><li>Prefix: {{.Meta.Prefix}}</li>
<li>VCS: {{ .Meta.VCS }}</li>
<li>RepoRoot: {{ .Meta.RepoRoot }}</li>
</ul>
</body>
</html>`
