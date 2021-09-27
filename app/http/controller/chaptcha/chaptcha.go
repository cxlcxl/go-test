package chaptcha

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/response"
	"net/http"
	"path"
	"strconv"
	"time"
)

type Captcha struct {
	Id      string `json:"id"`
	ImgUrl  string `json:"img_url"`
	Refresh string `json:"refresh"`
	Verify  string `json:"verify"`
}

// 生成验证码ID
func (c *Captcha) GenerateId(context *gin.Context) {
	// 设置验证码的数字长度（个数）
	var length = variable.ConfigYml.GetInt("Captcha.length")
	captchaId := captcha.NewLen(length)
	c.Id = captchaId
	c.ImgUrl = "/captcha/" + captchaId + ".png"
	c.Refresh = c.ImgUrl + "?reload=1"
	c.Verify = "/captcha/" + captchaId + "/这里替换为正确的验证码进行验证"
	response.Success(context, "验证码信息", c)
}

// 获取验证码图像
func (c *Captcha) GetImg(context *gin.Context) {
	captchaId := context.Param("captchaId")
	width := context.Query("width")
	height := context.Query("height")
	_, file := path.Split(context.Request.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if ext == "" || captchaId == "" {
		response.Fail(context, consts.CaptchaGetParamsInvalidCode, consts.CaptchaGetParamsInvalidMsg, "")
		return
	}

	if context.Query("reload") != "" {
		captcha.Reload(id)
	}

	context.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	context.Header("Pragma", "no-cache")
	context.Header("Expires", "0")

	var vBytes bytes.Buffer
	if ext == ".png" {
		context.Header("Content-Type", "image/png")
		captchaWidth, err := strconv.Atoi(width)
		if err != nil {
			captchaWidth = captcha.StdWidth
		}
		captchaHeight, err := strconv.Atoi(height)
		if err != nil {
			captchaWidth = captcha.StdHeight
		}
		_ = captcha.WriteImage(&vBytes, id, captchaWidth, captchaHeight)
		http.ServeContent(context.Writer, context.Request, id+ext, time.Time{}, bytes.NewReader(vBytes.Bytes()))
	}
}

// 校验验证码
func (c *Captcha) CheckCode(context *gin.Context) {
	captchaIdKey := variable.ConfigYml.GetString("Captcha.captchaId")
	captchaValueKey := variable.ConfigYml.GetString("Captcha.captchaValue")

	captchaId := context.Param(captchaIdKey)
	value := context.Param(captchaValueKey)

	if captchaId == "" || value == "" {
		response.Fail(context, consts.CaptchaCheckParamsInvalidCode, consts.CaptchaCheckParamsInvalidMsg, "")
		return
	}
	if captcha.VerifyString(captchaId, value) {
		response.Success(context, consts.CaptchaCheckOkMsg, "")
	} else {
		response.Fail(context, consts.CaptchaCheckFailCode, consts.CaptchaCheckFailMsg, "")
	}
}
