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
func CreateMysqlDB(database string) *gorm.DB {
	// 首字母大写
	database = strings.ToUpper(database[0:1]) + strings.ToLower(database[1:])
	var DbConn *gorm.DB
	lowerConnectName := strings.ToLower(database)
	if lowerConnectName == "api" {
		DbConn = apiDbConnect()
	} else if lowerConnectName == "data" {
		DbConn = dataDbConnect()
	} else {
		variable.ZapLog.Error("未知的数据库连接")
		return nil
	}
	return DbConn
}

func apiDbConnect() *gorm.DB {
	if variable.GormDbMysqlApi == nil {
		variable.GormDbMysqlApi = createDriver("Api")
	}
	return variable.GormDbMysqlApi
}

func dataDbConnect() *gorm.DB {
	if variable.GormDbMysqlData == nil {
		variable.GormDbMysqlData = createDriver("Data")
	}
	return variable.GormDbMysqlData
}

func createDriver(d string) (driver *gorm.DB) {
	var err error
	driver, err = gorm_v2.GetSqlDriver("Mysql", d, 1)
	if err != nil {
		variable.ZapLog.Error("创建 MYSQL 驱动失败，数据库为："+d, zap.Error(err))
		return nil
	}
	return
}
