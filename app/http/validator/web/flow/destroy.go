package flow

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	apiCtl "goskeleton/app/http/controller/web/apis"
	"goskeleton/app/http/validator"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type Destroy struct {
	Id float64 `form:"id"  json:"id" binding:"required,min=1"`
}

func (d Destroy) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&d); err != nil {
		response.ErrorParam(context, validator.Translate(err))
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(d, "", context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, consts.JsonMarshalFailed, "")
		return
	} else {
		(&apiCtl.AdApp{}).DeleteFlowConfig(extraAddBindDataContext)
	}
}
