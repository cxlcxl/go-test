package model

func UserGroupDB() *GroupModel {
	return &GroupModel{BaseModel: BaseModel{DB: UseDbConn("")}}
}

type GroupModel struct {
	BaseModel `json:"-"`
	BaseColumn
	TopModuleConfig string `gorm:"column:top_module_config" json:"top_module_config"`
	MenuConfig      string `gorm:"column:menu_config" json:"menu_config"`
	MenuRight       string `gorm:"column:menu_right" json:"menu_right"`
	TimeColumns
}

type BaseColumn struct {
	GroupId int64  `gorm:"column:group_id" json:"group_id"`
	Name    string `gorm:"column:name" json:"name"`
	Desc    string `gorm:"column:descrip" json:"desc"`
	Default int    `gorm:"column:default" json:"default"`
}

func (g *GroupModel) TableName() string {
	return "admin_group"
}

// GetGroup 查询全部
func (g *GroupModel) GetGroup(name string) (groups []*BaseColumn) {
	g.Table(g.TableName()).Where("`name` like ? and del = 0", "%"+name+"%").Order("group_id asc").Find(&groups)
	return
}
