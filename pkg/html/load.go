package html

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func LoadHTML(router *gin.Engine) {
	var files []string

	err := filepath.Walk("internal", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".tmpl") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	router.LoadHTMLFiles(files...)
}