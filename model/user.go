package model

import (
	"GenericModel/model_helper"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

//用户
type User struct {
	Id           int         `gorm:"PRIMARY_KEY"`                      //自增ID
	Name         string      `gorm:"type:VARCHAR(64);NOT NULL;UNIQUE"` //名称，不为空，唯一
	Description  string      `gorm:"size:128"`                         //用户描述
	DepartmentId int         `gorm:""`                                 //belongsTo foreign key：部门ID
	Department   *Department `gorm:""`                                 //BelongsTo：部门，每个用户必须有且只属于一个单位部门
	Profile      Profile     `gorm:""`                                 //HasOne：个人资料，每个用户必须有且只有一项个人资料
	Projects     []*Project  `gorm:""`                                 //HasMany：项目，用户负责的项目
	Roles        []*Role     `gorm:"many2many:role_user"`              //many2many，角色，每个用户可以有零个或多个角色，两表之间使用中间表'role_user'关联
	CreatedAt    time.Time   `gorm:""`                                 //创建时间
	UpdatedAt    time.Time   `gorm:""`                                 //更新时间
	DeletedAt    *time.Time  `gorm:""`                                 //删除时间
}

//-------------------------------------
// 业务
//-------------------------------------

/*
 * Description: 通过用户查询一行数据（不包含关联）
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *User) Query(db *gorm.DB) error {
	if m.Id == 0 {
		return errors.New("缺少用户ID")
	}
	return db.First(m).Error
}

/*
 * Description: 通过用户查询一行数据（只关联部门）
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *User) QueryWithDepartment(db *gorm.DB) error {
	if m.Id == 0 {
		return errors.New("缺少用户ID")
	}
	return db.Model(m).Preload("Department").First(m).Error
}

/*
 * Description: 通过用户查询一行数据（只关联个人资料）
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *User) QueryWithProfile(db *gorm.DB) error {
	if m.Id == 0 {
		return errors.New("缺少用户ID")
	}
	//return db.Model(m).Related(&m.Profile).First(m).Error//没有手机号
	return db.Model(m).Related(&m.Profile).Related(&m.Profile.Phones).First(m).Error
}

/*
 * Description: 通过用户查询一行数据（只关联用户项目）
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *User) QueryWithProject(db *gorm.DB) error {
	if m.Id == 0 {
		return errors.New("缺少用户ID")
	}
	return db.Model(m).Related(&m.Projects).First(m).Error
}

/*
 * Description: 通过用户查询一行数据（只关联用户角色）
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *User) QueryWithRole(db *gorm.DB) error {
	if m.Id == 0 {
		return errors.New("缺少用户ID")
	}
	return db.Model(m).Related(&m.Roles, "Roles").First(m).Error
}

/*
 * Description: 通过用户查询关联的角色
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *User) QueryRoles(db *gorm.DB) ([]*Role, error) {
	if m.Id == 0 {
		return nil, errors.New("缺少用户ID")
	}
	var roles []*Role
	err := db.Model(m).Association("Roles").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

/*
 * Description: 通过用户查询一行数据（关联所有）
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *User) QueryAll(db *gorm.DB) error {
	if m.Id == 0 {
		return errors.New("缺少用户ID")
	}
	return db.Model(m).
		Related(&m.Profile).
		Related(&m.Profile.Phones).
		Related(&m.Projects).
		Related(&m.Roles, "Roles").
		Preload("Department").First(m).Error
}

/*
 * Description: 查询用户列表（不包含关联）
 * Args:
 *  - db:
 *  - page:每几页
 *  - limit:每页显示几行
 * Returns:
 *  - []*User:用户列表
 *  - error:
 */
func (m *User) List(db *gorm.DB, page, limit int) ([]*User, error) {
	if page < 1 {
		page = 1
	}
	var users []*User
	err := db.Limit(limit).Offset((page - 1) * limit).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

/*
 * Description: 查询用户列表（包含所有关联）
 * Args:
 *  - db:
 *  - page:每几页
 *  - limit:每页显示几行
 * Returns:
 *  - []*User:用户列表
 *  - error:
 */
func (m *User) ListAll(db *gorm.DB, page, limit int) ([]*User, error) {
	if page < 1 {
		page = 1
	}
	var users []*User
	err := db.Limit(limit).Offset((page - 1) * limit).
		Preload("Roles").
		Preload("Profile").
		Preload("Profile.Phones").
		Preload("Projects").
		Preload("Department").
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

/*
 * Description: 创建一个用户（包含关联）
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *User) Create(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 在事务中做一些数据库操作 (这里应该使用 'tx' ，而不是 'db')
		if err := tx.Create(m).Error; err != nil {
			// 返回任意 err ，整个事务都会 rollback
			return err
		}
		// 返回 nil ，事务会 commit
		return nil
	})
}

/*
 * Description: 更新一个用户的角色
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *User) ModifyRole(db *gorm.DB, roles []*Role) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(m).Association("Roles").Replace(roles).Error; err != nil {
			return err
		}
		return nil
	})
}

/*
 * Description: 软删除一个用户
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *User) Delete(db *gorm.DB) error {
	if m.Id == 0 {
		return errors.New("缺少用户ID")
	}
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(m).Error; err != nil {
			return err
		}
		return nil
	})
}

/*
 * Description: 删除一个用户（包含关联）
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *User) Destroy(db *gorm.DB) error {
	if m.Id == 0 {
		return errors.New("缺少用户ID")
	}
	return db.Transaction(func(tx *gorm.DB) error {
		//个人资料
		profile := Profile{
			UserId: m.Id,
		}
		if err := profile.Destroy(tx); err != nil {
			return err
		}

		//用户项目
		project := &Project{
			UserId: m.Id,
		}
		err := project.DestroyByUserId(tx)
		if err != nil {
			return err
		}

		//用户角色
		if err := tx.Unscoped().Model(m).Association("Roles").Clear().Error; err != nil {
			return err
		}

		//用户
		if err := tx.Unscoped().Delete(m).Error; err != nil {
			return err
		}
		return nil
	})
}

//-------------------------------------
// 勾子
// 函数处理顺序：
//
// 创建：事务开始 -> BeforeSave -> BeforeCreate -> 业务 -> AfterCreate -> AfterSave -> 提交或回滚事务
// 更新：事务开始 -> BeforeSave -> BeforeUpdate -> 业务 -> AfterUpdate -> AfterSave -> 提交或回滚事务
// 删除：事务开始 -> BeforeDelete -> 业务 -> AfterDelete -> 提交或回滚事务
// 查询：预加载 -> BeforeDelete -> AfterFind
//
//-------------------------------------

/*
 * Description:
 * Args:
 *  - :
 * Returns:
 *  - :
 */
func (u *User) BeforeSave() (err error) {
	fmt.Println("BeforeSave:创建 或 更新 【前】 都需要执行这个函数")
	return nil
}

/*
 * Description:
 * Args:
 *  - :
 * Returns:
 *  - :
 */
func (u *User) AfterSave() (err error) {
	fmt.Println("AfterSave:创建 或 更新 【后】 都需要执行这个函数")
	return nil
}

/*
 * Description:
 * Args:
 *  - :
 * Returns:
 *  - :
 */
func (u *User) BeforeCreate() (err error) {
	if u.Name == "" {
		return errors.New("BeforeCreate:用户名称不能为空")
	}
	return nil
}

/*
 * Description:
 * Args:
 *  - :
 * Returns:
 *  - :
 */
func (u *User) AfterCreate() (err error) {
	u.Description = "AfterCreate:用户描述是在 'AfterCreate' 里更改的，并不改变数据库"
	return nil
}

/*
 * Description:
 * Args:
 *  - :
 * Returns:
 *  - :
 */
func (u *User) BeforeUpdate() (err error) {
	if u.Id == 0 {
		return errors.New("BeforeUpdate:更新用户时，ID不能为空")
	}
	return nil
}

/*
 * Description:
 * Args:
 *  - :
 * Returns:
 *  - :
 */
func (u *User) AfterUpdate() (err error) {
	u.Description = "AfterUpdate:用户描述是在 'AfterUpdate' 里更改的，并不改变数据库"
	return nil
}

/*
 * Description:
 * Args:
 *  - :
 * Returns:
 *  - :
 */
func (u *User) BeforeDelete() (err error) {
	if u.Id == 1 {
		return errors.New("BeforeDelete:不能删除用户ID为 1 的数据，试试删除ID为 2 的数据吧")
	}
	return nil
}

/*
 * Description:
 * Args:
 *  - :
 * Returns:
 *  - :
 */
func (u *User) AfterDelete() (err error) {
	if u.Id == 2 {
		return errors.New("AfterDelete:再挖个坑，删除ID为2的用户将会产生错误并回滚")
	}
	return nil
}

/*
 * Description:
 * Args:
 *  - :
 * Returns:
 *  - :
 */
func (u *User) AfterFind() (err error) {
	fmt.Println("AfterFind:每次查询后都会打印这行")
	return nil
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
func (m *User) DropTable(db *gorm.DB) error {
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
func (m *User) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(m).Error
}

/*
 * Description: 创建外键
 * Args:
 *  - db:
 * Returns:
 *  - error:
 */
func (m *User) CreateForeignKey(db *gorm.DB) error {
	return model_helper.CreateForeignKey(db, m, "department_id", "department(id)", "CASCADE", "CASCADE")
}
