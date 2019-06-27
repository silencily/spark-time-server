package controllers

import (
	"errors"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/mojocn/base64Captcha"
	"github.com/silencily/sparktime/core"
	"github.com/silencily/sparktime/models"
	"github.com/silencily/sparktime/services"
	"time"
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

func (c *SparkController) Post(ctx iris.Context) models.ResponseResult {
	result := models.ResponseResult{Code: 1, Message: "恭喜，火花发布成功！"}
	//校验验证码
	validCode := ctx.FormValue("validCode")
	sessionId := c.Session.ID()
	valid := base64Captcha.VerifyCaptcha(sessionId, validCode)
	if !valid {
		result.Code = 100 //权限错误
		result.Message = "验证码错误，请重试！"
		return result
	}

	file, info, err := ctx.FormFile("sparkImage")
	if err != nil {
		result.Code = 400 //参数错误
		result.Message = "发布失败，请稍后重试！"
	}
	imgName := info.Filename
	sparkContent := ctx.FormValue("sparkContent")

	spark := &models.Spark{
		Content:     sparkContent,
		CreatedTime: models.Time(time.Now()),
		ImgName:     imgName,
	}
	err = c.SparkService.Save(spark, file)
	if err != nil {
		getLogger().Errorf("spark保存失败:%s", err.Error())
		result.Code = 500 //内部错误
		result.Message = "发布失败，请稍后重试！"
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

func (c *SparkController) GetCaptcha() string {
	sessionId := c.Session.ID()
	captcha := core.GenerateCharacterCaptchaBase64Encoding(sessionId)
	return captcha
}
