package model

import (
	"GenericModel/model_helper"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

//用户的个人资料
type Profile struct {
	UserId   int       `gorm:"NOT NULL;UNIQUE"`                                 //外键：用户ID
	Gender   uint8     `gorm:"default:1"`                                       //性别：0未知，1男，2女
	Birthday time.Time `gorm:""`                                                //生日
	Phones   []*Phone  `gorm:"association_foreignkey:UserId;foreignkey:UserId"` //HasMany：手机，每个用户可以有零个或多个手机
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
func (m *Profile) DropTable(db *gorm.DB) error {
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
func (m *Profile) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(m).Error
}

/*
 * Description: 创建外键
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *Profile) CreateForeignKey(db *gorm.DB) error {
	return model_helper.CreateForeignKey(db, m, "user_id", "user(id)", "CASCADE", "CASCADE")
}

func (m *Profile) Destroy(tx *gorm.DB) error {
	if m.UserId == 0 {
		return errors.New("缺少用户ID")
	}
	var phone = &Phone{
		UserId: m.UserId,
	}
	err := phone.DestroyByUserId(tx)
	if err != nil {
		return err
	}
	return tx.Where("user_id = ?", m.UserId).Delete(m).Error
}
