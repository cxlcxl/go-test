package groups

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type Store struct {
	Name string `form:"name" json:"name" binding:"required,min=1"`
	Desc string `form:"desc" json:"desc"`
}

func (s Store) CheckParams(c *gin.Context) {
	if err := c.ShouldBind(&s); err != nil {
		response.ErrorParam(c, validator.Translate(err))
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(s, "", c)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(c, consts.JsonMarshalFailed, "")
	} else {
		(&web.Group{}).Store(extraAddBindDataContext)
	}
}
