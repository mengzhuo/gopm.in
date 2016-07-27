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
	prefix := []string{ctx.Request.URL.Scheme,
		cfg.Domain,
		"github",
		ctx.Param("owner"),
		repo,
	}
	meta := &MetaImport{
		VCS:    "git",
		Prefix: strings.Join(prefix, "/"),
	}
	fmt.Println(ctx.Request.RequestURI)
	ctx.JSON(http.StatusOK, meta)
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
