package main

import (
	"GenericModel/model"
	"github.com/jinzhu/gorm"
	"time"
)

func crateTables(db *gorm.DB) {
	var err error
	department := new(model.Department)
	if err = department.Migrate(db); err != nil {
		panic(err)
	}
	role := new(model.Role)
	if err = role.Migrate(db); err != nil {
		panic(err)
	}
	user := new(model.User)
	if err = user.Migrate(db); err != nil {
		panic(err)
	}
	profile := new(model.Profile)
	if err = profile.Migrate(db); err != nil {
		panic(err)
	}
	phone := new(model.Phone)
	if err = phone.Migrate(db); err != nil {
		panic(err)
	}
	project := new(model.Project)
	if err = project.Migrate(db); err != nil {
		panic(err)
	}
}

func dropTables(db *gorm.DB) {
	var err error
	project := new(model.Project)
	if err = project.DropTable(db); err != nil {
		panic(err)
	}
	phone := new(model.Phone)
	if err = phone.DropTable(db); err != nil {
		panic(err)
	}
	roleUser := new(model.RoleUser)
	if err = roleUser.DropTable(db); err != nil {
		panic(err)
	}
	profile := new(model.Profile)
	if err = profile.DropTable(db); err != nil {
		panic(err)
	}
	user := new(model.User)
	if err = user.DropTable(db); err != nil {
		panic(err)
	}
	role := new(model.Role)
	if err = role.DropTable(db); err != nil {
		panic(err)
	}
	department := new(model.Department)
	if err = department.DropTable(db); err != nil {
		panic(err)
	}
}
func createForeignKey(db *gorm.DB) {
	var err error
	user := new(model.User)
	if err = user.CreateForeignKey(db); err != nil {
		panic(err)
	}
	roleUser := new(model.RoleUser)
	if err = roleUser.CreateForeignKey(db); err != nil {
		panic(err)
	}
	project := new(model.Project)
	if err = project.CreateForeignKey(db); err != nil {
		panic(err)
	}
	profile := new(model.Profile)
	if err = profile.CreateForeignKey(db); err != nil {
		panic(err)
	}
	phone := new(model.Phone)
	if err = phone.CreateForeignKey(db); err != nil {
		panic(err)
	}
}

func demoData(db *gorm.DB) {
	var (
		departments = []*model.Department{
			{Name: "销售部"}, //1
			{Name: "生产部"}, //2
		}

		roles = []*model.Role{
			{Name: "销售代表"}, //1
			{Name: "销售主管"}, //2
			{Name: "销售总监"}, //3
			{Name: "工程师"},  //4
			{Name: "生产主管"}, //5
			{Name: "生产总监"}, //6
		}

		users = []*model.User{
			{Name: "赵一", DepartmentId: 1}, //1:1,2,3
			{Name: "孙二", DepartmentId: 2}, //2:4,5,6
			{Name: "张三", DepartmentId: 2}, //3:4,5,6
			{Name: "李四", DepartmentId: 1}, //4:1,2,3
			{Name: "王五", DepartmentId: 1}, //5:1,2,3
			{Name: "周六", DepartmentId: 2}, //6:4,5,6
			{Name: "李七", DepartmentId: 1}, //7:1,2,3
		}

		roleUsers = []*model.RoleUser{
			{1, 1},
			{3, 1},
			{5, 2},
			{6, 3},
			{4, 3},
			{5, 3},
			{3, 4},
			{1, 5},
			{6, 6},
			{4, 6},
			{2, 7},
		}

		projects = []*model.Project{
			{Name: "生产手机", UserId: 3},
			{Name: "生产电脑", UserId: 6},
			{Name: "销售手机", UserId: 1},
			{Name: "销售电脑", UserId: 4},
		}

		profiles = []*model.Profile{
			{UserId: 1, Gender: 2, Birthday: time.Now()},
			{UserId: 2, Birthday: time.Now()},
			{UserId: 3, Birthday: time.Now()},
			{UserId: 4, Gender: 1, Birthday: time.Now()},
			{UserId: 5, Gender: 0, Birthday: time.Now()},
			{UserId: 6, Birthday: time.Now()},
			{UserId: 7, Gender: 1, Birthday: time.Now()},
		}

		phones = []*model.Phone{
			{UserId: 1, Value: "13800000000"},
			{UserId: 1, Value: "13922223333"},
			{UserId: 2, Value: "13100004444"},
			{UserId: 3, Value: "13522223333"},
			{UserId: 3, Value: "13788889999"},
			{UserId: 3, Value: "15000003333"},
			{UserId: 4, Value: "15300223311"},
			{UserId: 5, Value: "13866664444"},
			{UserId: 6, Value: "15211223322"},
			{UserId: 6, Value: "15599997777"},
			{UserId: 7, Value: "16620000000"},
			{UserId: 6, Value: "13388885555"},
			{UserId: 6, Value: "13388885555"},
		}
	)

	for _, d := range departments {
		db.Create(d)
		time.Sleep(time.Millisecond * 300)
	}
	for _, d := range roles {
		db.Create(d)
		time.Sleep(time.Millisecond * 300)
	}
	for _, d := range users {
		db.Create(d)
		time.Sleep(time.Millisecond * 300)
	}
	for _, d := range roleUsers {
		db.Create(d)
		time.Sleep(time.Millisecond * 300)
	}
	for _, d := range projects {
		db.Create(d)
		time.Sleep(time.Millisecond * 300)
	}
	for _, d := range profiles {
		db.Create(d)
		time.Sleep(time.Millisecond * 300)
	}
	for _, d := range phones {
		db.Create(d)
		time.Sleep(time.Millisecond * 300)
	}
}
