package model

import (
	"GenericModel/model_helper"
	"github.com/jinzhu/gorm"
)

type RoleUser struct {
	RoleId int `gorm:"PRIMARY_KEY"`
	UserId int `gorm:"PRIMARY_KEY"`
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
func (m *RoleUser) DropTable(db *gorm.DB) error {
	//如果表存在，则删除
	return db.DropTableIfExists(m).Error
}

/*
 * Description: 创建外键
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *RoleUser) CreateForeignKey(db *gorm.DB) error {
	var err error
	if err = model_helper.CreateForeignKey(db, m, "role_id", "role(id)", "CASCADE", "CASCADE"); err != nil {
		return err
	}
	if err = model_helper.CreateForeignKey(db, m, "user_id", "user(id)", "CASCADE", "CASCADE"); err != nil {
		return err
	}
	return nil
}
