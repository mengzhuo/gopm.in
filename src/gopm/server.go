package gopm

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const gitRefsSuffix = ".git/info/refs?service=git-upload-pack"

var (
	repoTmpl *template.Template
	Router   *gin.Engine
)

func init() {
	var err error
	repoTmpl, err = template.New("repo").Parse(GoImportTmpl)
	if err != nil {
		log.Fatal(err)
	}
	Router = gin.Default()
	Router.SetHTMLTemplate(repoTmpl)

	Router.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.Status(200)
	})

	gb := Router.Group("/github")
	{
		gb.GET("/:owner/:repo", githubHandler)
		gb.GET("/:owner/:repo/.git/info/refs", githubHandler)
	}
}

func githubHandler(ctx *gin.Context) {
	repo := ctx.Param("repo")
	if strings.Count(repo, ".v") != 1 {
		err := fmt.Errorf("repo:%s is not supported", repo)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	Scheme := "http"
	if ctx.Request.TLS != nil {
		Scheme = "https"
	}
	ctx.Request.URL.Scheme = Scheme
	ctx.Request.URL.Host = ctx.Request.Host

	if ctx.Request.FormValue("go-get") == "1" {
		log.Print("go-get")
		delete(ctx.Request.URL.Query(), "go-get")
	}
	meta := &MetaImport{
		VCS:    "git",
		Prefix: ctx.Request.URL.String(),
	}
	fmt.Println(ctx.Request)
	ctx.HTML(http.StatusOK, "repo", meta)
}

type MetaImport struct {
	Prefix   string
	VCS      string
	RepoRoot string
}

const GoImportTmpl = `<html>
<head>
<meta content="{{.Prefix}} {{ .VCS }} {{ .RepoRoot }}" name="go-import" >
</head>
<body>
<ul><li>Prefix: {{.Prefix}}</li>
<li>VCS: {{ .VCS }}</li>
<li>RepoRoot: {{ .RepoRoot }}</li>
</ul>
</body>
</html>`
