package model

import (
	"GenericModel/model_helper"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

//用户的手机号码
type Phone struct {
	Id        int        `gorm:""`                              //主键，tag可以留空，gorm会自动添加自增主键
	UserId    int        `gorm:""`                              //外键：用户ID
	Profile   *Profile   `gorm:"association_foreignkey:UserId"` //用户这里可以通过手机号查询用户
	Value     string     `gorm:"size:11"`                       //手机号码
	CreatedAt time.Time  `gorm:""`                              //创建时间
	UpdatedAt time.Time  `gorm:""`                              //更新时间
	DeletedAt *time.Time `gorm:""`                              //删除时间
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
func (m *Phone) DropTable(db *gorm.DB) error {
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
func (m *Phone) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(m).Error
}

/*
 * Description: 创建外键
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *Phone) CreateForeignKey(db *gorm.DB) error {
	return model_helper.CreateForeignKey(db, m, "user_id", "user(id)", "CASCADE", "CASCADE")
}

func (m *Phone) Destroy(db *gorm.DB) error {
	if m.Id == 0 {
		return errors.New("缺少ID")
	}
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(m).Error; err != nil {
			return err
		}
		return nil
	})
}

func (m *Phone) DestroyByUserId(tx *gorm.DB) error {
	if m.UserId == 0 {
		return errors.New("缺少用户ID")
	}
	if err := tx.Unscoped().Where("user_id = ?", m.UserId).Delete(m).Error; err != nil {
		return err
	}
	return nil
}
