package views

import (
	"bytes"
	"html/template"
	"os"

	"github.com/kataras/iris/v12"

	"earth-assets/common"
)

func SendResponse(ctx iris.Context, statusCode int, errCode common.ErrCode, data interface{}) {
	ctx.StatusCode(statusCode)
	if statusCode != iris.StatusOK {
		return
	}

	resp := struct {
		status common.ErrCode
		data   interface{}
	}{
		status: errCode,
		data:   data,
	}
	ctx.JSON(resp)
}

// ParseTemplate generates html file
func ParseTemplate(data interface{}) error {
	templ, err := template.ParseFiles(common.IndexHTMLTempl)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = templ.Execute(buf, data); err != nil {
		return err
	}

	// generate html
	f, err := os.Create(common.IndexHTML)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write([]byte(buf.String())); err != nil {
		return err
	}

	return nil
}
