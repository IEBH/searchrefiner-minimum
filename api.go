package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hscells/cqr"
	"github.com/hscells/groove/combinator"
	gpipeline "github.com/hscells/groove/pipeline"
	"github.com/hscells/transmute"
	tpipeline "github.com/hscells/transmute/pipeline"
	log "github.com/sirupsen/logrus"
)

func handleTree(c *gin.Context) {
	fmt.Println("Hello, World!")
	rawQuery := c.PostForm("query")
	lang := c.PostForm("lang")
	username := c.PostForm("username")
	collect, err := strconv.ParseBool(c.PostForm("collect"))
	if err != nil {
		panic(err)
	}
	// Create relevant array of pmids {{{
	// Split string into array by newline
	pmids := strings.Fields(c.PostForm("pmids"))
	// Create array of integers
	var rel = []int{}
	for _, i := range pmids {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		rel = append(rel, j)
	}
	// Make combinator.Documents array
	relevant := make(combinator.Documents, len(rel))
	for i, r := range rel {
		relevant[i] = combinator.Document(r)
	}
	// }}}

	p := make(map[string]tpipeline.TransmutePipeline)
	p["medline"] = transmute.Medline2Cqr
	p["pubmed"] = transmute.Pubmed2Cqr

	compiler := p["medline"]
	if v, ok := p[lang]; ok {
		compiler = v
	} else {
		lang = "medline"
	}

	cq, err := compiler.Execute(rawQuery)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	repr, err := cq.Representation()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return

	}

	var root combinator.LogicalTree
	root, err = combinator.NewShallowLogicalTree(gpipeline.NewQuery("searchrefiner", "0", repr.(cqr.CommonQueryRepresentation)), entrez, relevant)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	t := buildTree(root.Root, relevant...)

	t.NumRel = len(relevant)
	switch r := root.Root.(type) {
	case combinator.Combinator:
		t.NumRelRet = int(r.R)
	case combinator.Atom:
		t.NumRelRet = int(r.R)
	}

	var numRet int64
	if len(t.Nodes) > 0 {
		numRet = int64(t.Nodes[0].Value)
	}

	if collect {
		log.Infof(fmt.Sprintf("[username=%s][query=%s][lang=%s][pmids=%v][numrel=%d][numret=%d][numrelret=%d]", username, rawQuery, lang, relevant, t.NumRel, numRet, t.NumRelRet))
	}

	c.JSON(200, t)
}
