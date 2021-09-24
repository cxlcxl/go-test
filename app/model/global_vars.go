package model

const (
	CheckPass     = 1
	CheckNotPass  = -1
	CheckIng      = 2
	OpenStatus    = 1
	CloseStatus   = 1
	OperatorUser  = 1
	AdsUser       = 2
	DeveloperUser = 3
	GDTUser       = 4
	SPMUser       = 5
	IDTUser       = 6
)

var (
	ConfTypes = map[int]string{1: "全局配置", 2: "定向配置"}
	AppTypes  = map[int]string{3: "联盟流量", 1: "自有流量", 4: "小游戏流量", 5: "休闲游戏变现"}
	UserTypes = map[int]string{OperatorUser: "营用户", AdsUser: "广告主用户", DeveloperUser: "开发者用户", GDTUser: "点通用户", SPMUser: "投放用户", IDTUser: "IDT用户"}
)
