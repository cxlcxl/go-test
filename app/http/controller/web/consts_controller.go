package web

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/model"
	"goskeleton/app/utils/response"
	"strings"
)

type GlobalConst struct {
}

// GetGlobalConst 公共常量
func (cs *GlobalConst) GetGlobalConst(c *gin.Context) {
	code := c.Param("constCode")
	var data interface{}
	switch strings.ToLower(code) {
	case "M-AD-SUB-TYPE":
		data = model.MAdSubType
	case "M-AD-TYPE":
		data = model.MAdType
	case "PLATFORM":
		data = model.MPlatformDesc
	case "AD-TYPES":
		data = model.AppTypes
	case "M-AD-SUB-TYPE-DESC":
		data = model.MAdSubTypeDesc
	case "M-AD-POS-TYPE":
		data = model.MAdPosType
	case "M-AD-POS-TYPE-NAME":
		data = model.MAdPosTypeName
	case "M-CUSTOM-SUB-TYPE":
		data = model.MCustomSubType
	case "M-EMBED-SUB-TYPE":
		data = model.MEmbedSubType
	case "M-NEW-EMBED-SUB-TYPE":
		data = model.MNewEmbedSubType
	case "M-TEMPLATE-SUB-TYPE":
		data = model.MTemplateSubType
	case "M-EMBED-SIZE":
		data = model.MEmbedSize
	case "M-BANNER-SIZE":
		data = model.MBannerSize
	case "M-NEW-EMBED-SIZE":
		data = model.MNewEmbedSize
	case "USER-TYPE":
		data = model.UserType
	case "USER-CHECK-TYPE":
		data = model.UserCheckType
	default:
		response.Fail(c, consts.CurdSelectFailCode, "不存在请求的常量", "")
		return
	}
	response.Success(c, consts.CurdStatusOkMsg, data)
}
