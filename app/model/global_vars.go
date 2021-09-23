package model

const (
	CheckPass    = 1
	CheckNotPass = -1
	CheckIng     = 2
	OpenStatus   = 1
	CloseStatus  = 1
)

var (
	ConfTypes = map[int]string{1: "全局配置", 2: "定向配置"}
	AppTypes  = map[int]string{3: "联盟流量", 1: "自有流量", 4: "小游戏流量", 5: "休闲游戏变现"}
)
