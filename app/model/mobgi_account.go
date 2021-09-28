package model

import (
	"goskeleton/app/global/variable"
	"time"
)

func MobgiAccountDB() *MobgiAccountModel {
	return &MobgiAccountModel{BaseModel: BaseModel{DB: UseDbConn("")}}
}

type MobgiAccountModel struct {
	BaseModel `json:"-"`
	AccountBaseColumns
	CreditCode        string `gorm:"column:credit_code" json:"credit_code"`
	LegalPerson       string `gorm:"column:legal_person" json:"legal_person"`               // 法人姓名
	CompanyMobile     string `gorm:"column:company_mobile" json:"company_mobile"`           // 联系人电话
	BusinessLicense   string `gorm:"column:business_license" json:"business_license"`       // 营业执照
	Address           string `gorm:"column:address" json:"address"`                         // 注册地址
	Contact           string `gorm:"column:contact" json:"contact"`                         // 商务联系人
	Mobile            string `gorm:"column:mobile" json:"mobile"`                           // 联系人电话
	Email             string `gorm:"column:email" json:"email"`                             // 联系人邮箱
	CompanyAddress    string `gorm:"column:company_address" json:"company_address"`         // 通讯地址
	BankName          string `gorm:"column:bank_name" json:"bank_name"`                     // 对公银行名称
	City              string `gorm:"column:city" json:"city"`                               // 对公开户省市
	BankBranchName    string `gorm:"column:bank_branch_name" json:"bank_branch_name"`       // 对公开户支行
	BankAccountNumber string `gorm:"column:bank_account_number" json:"bank_account_number"` // 对公银行账户
	InvoiceType       int    `gorm:"column:invoice_type" json:"invoice_type"`               // 发票种类：1普通发票；2增值税专用发票
	ContentType       int    `gorm:"column:content_type" json:"content_type"`               // 发票内容：1广告费；2技术服务费
	TimeColumns
}

type AccountBaseColumns struct {
	Id          int64  `json:"id" gorm:"column:id"`
	CompanyName string `json:"company_name" gorm:"column:company_name"` // 公司名称
}

// TableName 表名
func (m *MobgiAccountModel) TableName() string {
	return "admin_mobgi_account"
}

// Store 新增
func (m *MobgiAccountModel) Store(title, des, content string, userId int) bool {
	sql := "INSERT INTO news(user_id,title,des,content,created_at) VALUES (?,?,?,?,?)"
	if m.Exec(sql, userId, title, des, content, time.Now().Format(variable.DateFormat)).RowsAffected > 0 {
		return true
	}
	return false
}

// GetAll ...
func (m *MobgiAccountModel) GetAll() (info []*AccountBaseColumns) {
	m.Find(&info)
	return
}
