package model

import (
	"GenericModel/model_helper"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

//项目，假设每个项目只能由一个用户管理
type Project struct {
	Id        int        `gorm:""`         //主键
	Name      string     `gorm:"NOT NULL"` //项目名称
	UserId    int        `gorm:"NOT NULL"` //外键：用户ID
	User      *User      `gorm:""`         //用户这里可以通过手机号查询用户
	CreatedAt time.Time  `gorm:""`         //创建时间
	UpdatedAt time.Time  `gorm:""`         //更新时间
	DeletedAt *time.Time `gorm:""`         //删除时间
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
func (m *Project) DropTable(db *gorm.DB) error {
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
func (m *Project) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(m).Error
}

/*
 * Description: 创建外键
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *Project) CreateForeignKey(db *gorm.DB) error {
	return model_helper.CreateForeignKey(db, m, "user_id", "user(id)", "CASCADE", "CASCADE")
}

func (m *Project) Destroy(db *gorm.DB) error {
	return nil
}

func (m *Project) DestroyByUserId(tx *gorm.DB) error {
	if m.UserId == 0 {
		return errors.New("缺少用户ID")
	}
	return tx.Where("user_id = ?", m.UserId).Delete(m).Error
}
