package model

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"goskeleton/app/model/tool"
	"goskeleton/app/utils/gorm_v2"
	"strings"
)

type BaseModel struct {
	*gorm.DB `gorm:"-" json:"-"`
}

type IdColumns struct {
	Id int64 `gorm:"column:id" json:"id"`
}

type BaseColumns struct {
	Id        int64          `gorm:"primarykey" json:"id"`
	CreatedAt tool.LocalTime `json:"created_at"`
	UpdatedAt tool.LocalTime `json:"updated_at"`
}

type TimeColumns struct {
	CreatedAt tool.LocalTime `json:"create_time" gorm:"column:create_time"`
	UpdatedAt tool.LocalTime `json:"update_time" gorm:"column:update_time"`
}

func UseDbConn(sqlType string) *gorm.DB {
	var db *gorm.DB
	sqlType = strings.Trim(sqlType, " ")
	if sqlType == "" {
		sqlType = variable.ConfigGormv2Yml.GetString("Gormv2.UseDbType")
	}
	switch strings.ToLower(sqlType) {
	case "mysql":
		if variable.GormDbMysql == nil {
			variable.ZapLog.Fatal(fmt.Sprintf(my_errors.ErrorsGormNotInitGlobalPointer, sqlType, sqlType))
		}
		db = variable.GormDbMysql
	case "sqlserver":
		if variable.GormDbSqlserver == nil {
			variable.ZapLog.Fatal(fmt.Sprintf(my_errors.ErrorsGormNotInitGlobalPointer, sqlType, sqlType))
		}
		db = variable.GormDbSqlserver
	case "postgres", "postgre", "postgresql":
		if variable.GormDbPostgreSql == nil {
			variable.ZapLog.Fatal(fmt.Sprintf(my_errors.ErrorsGormNotInitGlobalPointer, sqlType, sqlType))
		}
		db = variable.GormDbPostgreSql
	default:
		variable.ZapLog.Error(my_errors.ErrorsDbDriverNotExists + sqlType)
	}
	return db
}

// CreateMysqlDB 创建一个连接
//func CreateMysqlDB(database string) *gorm.DB {
//	sqlType := "Mysql"
//	if variable.GormDbMysqlData == nil {
//		// 首字母大写
//		database = strings.ToUpper(database[0:1]) + strings.ToLower(database[1:])
//		var err error
//		variable.GormDbMysqlData, err = gorm_v2.GetSqlDriver(sqlType, database, 1)
//		if err != nil {
//			variable.ZapLog.Error(my_errors.ErrorsDialectorDbInitFail+sqlType, zap.Error(err))
//		}
//	}
//	return variable.GormDbMysqlData
//}

// CreateMysqlDB 创建一个连接
func CreateMysqlDB(database string) *gorm.DB {
	// 首字母大写
	database = strings.ToUpper(database[0:1]) + strings.ToLower(database[1:])
	var err error
	sqlType := "Mysql"
	DbConn := variable.GormDbMysqlData
	switch strings.ToLower(database) {
	case "api":
		DbConn = variable.GormDbMysqlApi
	default:
		DbConn = variable.GormDbMysqlData
	}
	if DbConn == nil {
		DbConn, err = gorm_v2.GetSqlDriver(sqlType, database, 1)
		if err != nil {
			variable.ZapLog.Error(my_errors.ErrorsDialectorDbInitFail+sqlType, zap.Error(err))
		}
	}
	return DbConn
}
