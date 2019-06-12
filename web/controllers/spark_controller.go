package controllers

import (
	"errors"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/silencily/sparktime/core"
	"github.com/silencily/sparktime/services"
)

type SparkController struct {
	SparkService services.SparkService
}

func getLogger() *golog.Logger {
	return core.GetLogger("SparkController")
}

func (c *SparkController) Get() map[string]interface{} {
	result := map[string]interface{}{
		"name": "foo",
		"time": 2,
	}
	return result
}

func (c *SparkController) GetImgBy(docId string) mvc.Result {
	result := mvc.Response{}
	img, contentType := c.SparkService.GetImg(docId)
	if img != nil {
		result.ContentType = contentType
		result.Content = img
	} else {
		getLogger().Errorf("DocId: %s not found.", docId)
		result.Err = errors.New("not found")
		result.Code = iris.StatusNotFound
	}
	return result

}