package main

import (
	"bytes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"unoconv/models"
	"unoconv/until"
)

const ErrorCode = 70001

var (
	uno *until.Unoconv
)

func init() {
	uno = until.InitUnoconv()
}

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "**** time=${time_rfc3339} , host=${host} , method=${method} ," +
			" uri=${uri} , form_path=${form:path}, form_notify=${form:notify} , status=${status} ***\n",
	}))
	extDict := map[string]string{"txt": "1", "doc": "1", "docx": "1", "pdf": "1", "rtf": "1", "xls": "1", "xlsx": "1", "ppt": "1", "pptx": "1"}

	e.GET("/", func(c echo.Context) error {
		msg := models.MsgReturn{Code: ErrorCode, Content: "path不能为空"}
		return c.JSON(http.StatusOK, msg)
	})
	e.POST("/unoconv", func(c echo.Context) error {
		path := c.FormValue("path")
		notify := c.FormValue("notify")
		ext, _ := until.GetFileExt(path)
		msg := models.MsgReturn{}
		if len(path) == 0 || len(notify) == 0 || extDict[ext] != "1" {
			msg = models.MsgReturn{Code: ErrorCode, Content: "参数错误"}
			return c.JSON(http.StatusOK, msg)
		}
		buf := bytes.NewBufferString("")
		go uno.Convert("pdf", buf, notify, path)
		return c.JSON(http.StatusOK, msg)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
