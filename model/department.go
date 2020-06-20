package model

import (
	"github.com/jinzhu/gorm"
)

//用户所属部门，比如生产部门，销售部门
type Department struct {
	Id   int    `gorm:"PRIMARY_KEY"`
	Name string `gorm:"size:64"`
	User *[]User
}

//-------------------------------------
// 数据表操作
//-------------------------------------

/*
 * Description: 删除表
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *Department) DropTable(db *gorm.DB) error {
	//如果表存在，则删除
	return db.DropTableIfExists(m).Error
}

/*
 * Description: 迁移
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *Department) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(m).Error
}
