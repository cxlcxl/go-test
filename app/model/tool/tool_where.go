package tool

import (
	"fmt"
	"strings"
)

type WhereQuery struct {
	Filter      bool
	QuerySql    string        // where sql
	Queries     []string      // where 条件
	QueryParams []interface{} // ? 的参数
}

// GenerateWhere 生成 where 条件的 SQL
// filterEmpty 是否排除空参数
/**
wheres := map[string]interface{}{
	"username": []string{"like", "admin%|editor"}, // or查询 不加任何 % 为默认 %admin%
	"username": []string{"like", "admin%"}, // 不加任何 % 为默认 %admin%
	"email":    []string{"in", "admin,dd,gg,ee"},
	"or": map[string]interface{}{
		"username": []string{"like", "admin"},
		"age":      8,
	},
	"age": 6,
	"updated_at": "null",
	"created_at": "not null",
}
*/
func (w *WhereQuery) GenerateWhere(wheres map[string]interface{}) *WhereQuery {
	if len(wheres) == 0 {
		return nil
	}
	//sqlParams := make([]string, 0)
	for field, val := range wheres {
		switch val.(type) {
		case string:
			if w.Filter && len(val.(string)) == 0 {
				continue
			}
			w.stringWhere(field, val.(string))
		case int, int8, int16, int32, int64, float64, float32:
			w.Queries = append(w.Queries, fmt.Sprintf("`%s`=?", field))
			w.QueryParams = append(w.QueryParams, val)
		case []string:
			w.mapWhere(field, val.([]string))
		case map[string]interface{}:
			w.orWhere(val.(map[string]interface{}))
		default:
			panic("unknown where type, key: " + field)
		}
	}
	w.QuerySql = strings.Join(w.Queries, " AND ")
	return w
}

// mapWhere
func (w *WhereQuery) mapWhere(field string, whereSlice []string) {
	if len(whereSlice) != 2 && len(whereSlice) != 3 {
		panic("WHERE 子条件组装有误")
	}
	if w.Filter && len(whereSlice[1]) == 0 {
		return
	}
	switch strings.ToLower(whereSlice[0]) {
	case "like":
		if strings.Contains(whereSlice[1], "|") {
			w.orValWhere(field, whereSlice[1], "like")
		} else {
			w.Queries = append(w.Queries, fmt.Sprintf("`%s` LIKE ?", field))
			w.QueryParams = append(w.QueryParams, w.likeFormat(whereSlice[1]))
		}
	case "lt", "<":
		w.Queries = append(w.Queries, fmt.Sprintf("`%s` < ?", field))
		w.QueryParams = append(w.QueryParams, whereSlice[1])
	case "elt", "<=":
		w.Queries = append(w.Queries, fmt.Sprintf("`%s` <= ?", field))
		w.QueryParams = append(w.QueryParams, whereSlice[1])
	case "gt", ">":
		w.Queries = append(w.Queries, fmt.Sprintf("`%s` > ?", field))
		w.QueryParams = append(w.QueryParams, whereSlice[1])
	case "egt", ">=":
		w.Queries = append(w.Queries, fmt.Sprintf("`%s` >= ?", field))
		w.QueryParams = append(w.QueryParams, whereSlice[1])
	case "neq", "!=", "<>":
		w.Queries = append(w.Queries, fmt.Sprintf("`%s` != ?", field))
		w.QueryParams = append(w.QueryParams, whereSlice[1])
	case "in":
		w.Queries = append(w.Queries, fmt.Sprintf("`%s` IN (?)", field))
		w.QueryParams = append(w.QueryParams, "'"+strings.Join(strings.Split(whereSlice[1], ","), "','")+"'")
	case "between":
		w.Queries = append(w.Queries, fmt.Sprintf("`%s` BETWEEN ? AND ?", field))
		w.QueryParams = append(w.QueryParams, whereSlice[1], whereSlice[2])
	default:
		panic("unknown key type: " + whereSlice[0])
	}

	return
}

// orWhere
func (w *WhereQuery) orWhere(wheres map[string]interface{}) {
	whereMap := make([]string, 0)
	for field, val := range wheres {
		switch val.(type) {
		case string:
			if w.Filter && len(val.(string)) == 0 {
				continue
			}
			whereMap = append(whereMap, fmt.Sprintf("`%s`=?", field))
			w.QueryParams = append(w.QueryParams, val)
		case int, int8, int16, int32, int64, float64, float32:
			whereMap = append(whereMap, fmt.Sprintf("`%s`=?", field))
			w.QueryParams = append(w.QueryParams, val)
		case []string:
			w.mapWhere(field, val.([]string))
		default:
			panic("unknown where type, key: " + field)
		}
	}
	w.Queries = append(w.Queries, "("+strings.Join(whereMap, " OR ")+")")
	return
}

// orValWhere
func (w *WhereQuery) orValWhere(field, val, calc string) {
	values := strings.Split(val, "|")
	whereMap := make([]string, len(values))
	for k, v := range values {
		if strings.ToLower(calc) == "like" {
			v = w.likeFormat(v)
		}
		whereMap[k] = fmt.Sprintf("`%s` %s ?", field, calc)
		w.QueryParams = append(w.QueryParams, v)
	}
	w.Queries = append(w.Queries, "("+strings.Join(whereMap, " OR ")+")")
	return
}

// likeFormat
func (w *WhereQuery) likeFormat(s string) string {
	if hasBuf := strings.Contains(s, "%"); !hasBuf {
		s = "%" + s + "%"
	}
	return s
}

// stringWhere 字符串类型的值 = 号赋值情况
func (w *WhereQuery) stringWhere(field, val string) {
	if strings.Contains(val, "|") {
		w.orValWhere(field, val, "=")
	} else if strings.ToLower(val) == "null" {
		w.Queries = append(w.Queries, fmt.Sprintf("`%s` IS NULL", field))
	} else if strings.ToLower(val) == "not null" {
		w.Queries = append(w.Queries, fmt.Sprintf("`%s` IS NOT NULL", field))
	} else {
		w.Queries = append(w.Queries, fmt.Sprintf("`%s`=?", field))
		w.QueryParams = append(w.QueryParams, val)
	}
	return
}
