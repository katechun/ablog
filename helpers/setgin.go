package helpers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"path/filepath"
	"wblog/helpers"
)

func SetTemplate(engine *gin.Engine) {

	funcMap := template.FuncMap{
		"dateFormat": helpers.DateFormat,
		"substring":  helpers.Substring,
		"isOdd":      helpers.IsOdd,
		"isEven":     helpers.IsEven,
		"truncate":   helpers.Truncate,
		"add":        helpers.Add,
		"minus":      helpers.Minus,
		"listtag":    helpers.ListTag,
	}

	engine.SetFuncMap(funcMap)
	Logger.Println(filepath.Join(GetOneValue("app","htmldir"), "/**/*"))
	engine.LoadHTMLGlob(filepath.Join(GetOneValue("app","htmldir"), "/**/*"))
}