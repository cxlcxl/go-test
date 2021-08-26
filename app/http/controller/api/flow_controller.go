package api

import (
	"goskeleton/app/service/fs"
	"log"

	"github.com/gin-gonic/gin"
)

type Flow struct{}

func (f *Flow) HandleFlowExposure(context *gin.Context) {
	// 查询有设置曝光量的广告商配置项
	// adsRels := model.CreateFlowAdsRelFactory("").FetchRows("SELECT * FROM flow_ads_rel WHERE exposure_num > 0 AND conf_type IN (1, 2)")
	// fmt.Printf("%#v", adsRels)
	err := fs.CreateFsFactory().AccessToken().SendMsg("xiaoliang.cheng", "测试GO发送飞书消息")
	if err != nil {
		log.Println(err)
	}

}
