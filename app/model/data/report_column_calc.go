package data

var (
	BaseDims = []string{"app_key", "year_months", "days", "hours", "platform", "ad_type", "channel_gid", "ads_gid", "ads_id"}
)

// GetFilterFields 获取过滤字段
func GetFilterFields() []string {
	fields := []string{"ssp_id", "ads_gid", "ads_id", "pos_key", "app_version", "sdk_version", "sys_version"}
	return FieldMerge(BaseDims, fields)
}

// FieldMerge 字段合并
func FieldMerge(a1, a2 []string) []string {
	s := make(map[string]int, 0)
	for _, a := range a1 {
		s[a] = 1
	}
	for _, a := range a2 {
		if _, ok := s[a]; !ok {
			a1 = append(a1, a)
		}
	}
	return a1
}
