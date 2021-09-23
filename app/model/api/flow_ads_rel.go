package model_api

import (
	"fmt"
	"goskeleton/app/global/variable"
	"goskeleton/app/model"
	"time"

	"go.uber.org/zap"
)

func CreateFlowAdsRelFactory(sqlType string) *FlowAdsRelModel {
	return &FlowAdsRelModel{BaseModel: model.BaseModel{DB: model.UseDbConn(sqlType)}}
}

type FlowAdsRelModel struct {
	model.BaseModel `json:"-"`
	FlowId          int     `gorm:"column:flow_id" json:"flow_id"` // 流量配置id
	AppKey          string  `gorm:"column:app_key" json:"app_key"`
	AdType          int     `gorm:"column:ad_type" json:"ad_type"`             // 广告类型 1视频2插页
	ConfType        int     `gorm:"column:conf_type" json:"conf_type"`         // 1一般广告商2优先广告商3DSP配置广告商
	AdsId           string  `gorm:"column:ads_id" json:"ads_id"`               // 广告商id
	Position        int     `json:"position"`                                  // 位置
	LimitNum        int     `gorm:"column:limit_num" json:"limit_num"`         // 限制次数
	ExposureNum     int     `gorm:"column:exposure_num" json:"exposure_num"`   // 曝光量限制
	ReqLimitNum     int     `gorm:"column:req_limit_num" json:"req_limit_num"` // 请求次数限制
	LimitTime       int     `gorm:"column:limit_time" json:"limit_time"`       // 限制请求间隔时间
	Weight          float64 `json:"weight"`                                    // 权重
	Del             int     `json:"del"`                                       // 删除标志，0未删除1删除
}

// 表名
func (f *FlowAdsRelModel) TableName() string {
	return "flow_ads_rel"
}

// 查询
func (f *FlowAdsRelModel) FetchRows(sql string, query ...interface{}) *FlowAdsRelModel {
	// result := f.Select(sql, query...).Find()
	result := f.Select(sql, query...).Find(f)
	if result.Error == nil {
		return f
	} else {
		fmt.Println("err0r", result.Error)
		return nil
	}
}

// 查询单条数据
func (f *FlowAdsRelModel) FetchOne(sql string, query ...interface{}) *FlowAdsRelModel {
	result := f.Raw(sql, query...).First(f)
	if result.Error == nil {
		return f
	} else {
		variable.ZapLog.Error("查询出错:", zap.Error(result.Error))
	}
	return nil
}

// 用户刷新token
func (u *FlowAdsRelModel) OauthRefreshToken(userId, expiresAt int64, oldToken, newToken, clientIp string) bool {
	sql := "UPDATE   tb_oauth_access_tokens   SET  token=? ,expires_at=?,client_ip=?,updated_at=NOW(),action_name='refresh'  WHERE   fr_user_id=? AND token=?"
	if u.Exec(sql, newToken, time.Unix(expiresAt, 0).Format(variable.DateFormat), clientIp, userId, oldToken).Error == nil {
		return true
	}
	return false
}
