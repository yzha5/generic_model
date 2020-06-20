package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

//用户角色
type Role struct {
	Id        int        `gorm:"PRIMARY_KEY"`         //主键
	Name      string     `gorm:"size:64;NOT NULL"`    //角色名称，不能为空，但可以重复
	Users     []*User    `gorm:"many2many:role_user"` //many2many：
	CreatedAt time.Time  `gorm:""`                    //创建时间
	UpdatedAt time.Time  `gorm:""`                    //更新时间
	DeletedAt *time.Time `gorm:""`                    //删除时间
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
func (m *Role) DropTable(db *gorm.DB) error {
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
func (m *Role) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(m).Error
}
