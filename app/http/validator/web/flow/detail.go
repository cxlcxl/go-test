package flow

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	apiCtl "goskeleton/app/http/controller/web/api"
	"goskeleton/app/http/validator"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type DetailFlow struct {
	FlowId   float64 `form:"flow_id"  json:"flow_id" binding:"required,min=1"`
	ConfType float64 `form:"conf_type"  json:"conf_type" binding:"required"`
	AppKey   string  `form:"app_key"  json:"app_key" binding:"required"`
}

func (f DetailFlow) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&f); err != nil {
		response.ErrorParam(context, validator.Translate(err))
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(f, "", context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, consts.JsonMarshalFailed, "")
		return
	} else {
		(&apiCtl.AdApp{}).FlowConfigDetail(extraAddBindDataContext)
	}
}
