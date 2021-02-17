package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type Member struct{
  leader string
  sub string
}

type Contents struct{
  header string
  content string
}

type Document struct{
  name string
  member Member
  contents Contents
}

func (d Document) dict() string{
  return d.contents.content
}
func main() {
  router := gin.Default()
  var member = Member{"hirata","wada"}
  var contents = Contents{"This is Test","testing is important."}
  var doc = Document{"go-test",member,contents}
  router.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, doc.dict())
  })
  router.Run(":8080")
}