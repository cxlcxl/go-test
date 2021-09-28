package web

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/model/tool"
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
		data = tool.MAdSubType
	case "M-AD-TYPE":
		data = tool.MAdType
	case "PLATFORM":
		data = tool.MPlatformDesc
	case "AD-TYPES":
		data = tool.AppTypes
	case "M-AD-SUB-TYPE-DESC":
		data = tool.MAdSubTypeDesc
	case "M-AD-POS-TYPE":
		data = tool.MAdPosType
	case "M-AD-POS-TYPE-NAME":
		data = tool.MAdPosTypeName
	case "M-CUSTOM-SUB-TYPE":
		data = tool.MCustomSubType
	case "M-EMBED-SUB-TYPE":
		data = tool.MEmbedSubType
	case "M-NEW-EMBED-SUB-TYPE":
		data = tool.MNewEmbedSubType
	case "M-TEMPLATE-SUB-TYPE":
		data = tool.MTemplateSubType
	case "M-EMBED-SIZE":
		data = tool.MEmbedSize
	case "M-BANNER-SIZE":
		data = tool.MBannerSize
	case "M-NEW-EMBED-SIZE":
		data = tool.MNewEmbedSize
	case "USER-TYPE":
		data = tool.UserType
	case "USER-CHECK-TYPE":
		data = tool.UserCheckType
	default:
		response.Fail(c, consts.CurdSelectFailCode, "不存在请求的常量", "")
		return
	}
	response.Success(c, consts.CurdStatusOkMsg, data)
}
