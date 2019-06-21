package controllers

import (
	"errors"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/silencily/sparktime/core"
	"github.com/silencily/sparktime/models"
	"github.com/silencily/sparktime/services"
)

type SparkController struct {
	SparkService services.SparkService
	Session      *sessions.Session
}

func getLogger() *golog.Logger {
	return core.GetLogger("SparkController")
}

func (c *SparkController) Get() []models.Spark {
	sparks := c.SparkService.List()
	return sparks
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

func (c *SparkController) GetCaptcha() string {
	sessionId := c.Session.ID()
	captcha := core.GenerateCharacterCaptchaBase64Encoding(sessionId)
	return captcha
}
