package core

import "github.com/mojocn/base64Captcha"

//生成base64编码的图片验证码（算数）
func GenerateCharacterCaptchaBase64Encoding(idKey string) string {
	var config = base64Captcha.ConfigCharacter{
		Height: 50,
		Width:  100,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeArithmetic,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsUseSimpleFont:    true,
		IsShowHollowLine:   true,
		IsShowNoiseDot:     true,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
	}
	_, cap := base64Captcha.GenerateCaptcha(idKey, config)

	base64string := base64Captcha.CaptchaWriteToBase64Encoding(cap)

	return base64string
}
