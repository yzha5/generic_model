package model

import (
	"GenericModel/config"
	"GenericModel/source"
	"fmt"
	"testing"
	"time"
)

var (
	db   = source.NewDB()
	user = &User{Id: 1}
)

func TestUser_Query(t *testing.T) {
	err := user.Query(db)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}

func TestUser_QueryWithDepartment(t *testing.T) {
	err := user.QueryWithDepartment(db)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
func TestUser_QueryWithProfile(t *testing.T) {
	err := user.QueryWithProfile(db)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	fmt.Printf("%+v\n", user.Profile)
	fmt.Printf("%+v\n", user.Profile.Phones)
}

func TestUser_QueryWithProject(t *testing.T) {
	err := user.QueryWithProject(db)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	fmt.Printf("%+v\n", user.Projects[0])
}

func TestUser_QueryWithRole(t *testing.T) {
	err := user.QueryWithRole(db)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	fmt.Printf("%+v\n", user.Roles)
	fmt.Printf("%+v\n", user.Roles[0])
}

func TestUser_QueryAll(t *testing.T) {
	err := user.QueryAll(db)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	fmt.Printf("%+v\n", user.Department)
	fmt.Printf("%+v\n", user.Projects)
	fmt.Printf("%+v\n", user.Projects[0])
	fmt.Printf("%+v\n", user.Profile)
	fmt.Printf("%+v\n", user.Profile.Phones)
	fmt.Printf("%+v\n", user.Profile.Phones[0])
	fmt.Printf("%+v\n", user.Roles)
	fmt.Printf("%+v\n", user.Roles[0])
}

func TestUser_QueryRole(t *testing.T) {
	roles, err := user.QueryRoles(db)
	if err != nil {
		panic(err)
	}
	for _, role := range roles {
		fmt.Println(role)
	}
}

func TestUser_List(t *testing.T) {
	users, err := user.List(db, 1, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
}

func TestUser_ListAll(t *testing.T) {
	users, err := user.ListAll(db, 1, 10)
	if err != nil {
		panic(err)
	}
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
}

func TestUser_Create(t *testing.T) {
	birthday, _ := time.ParseInLocation(config.TimeTemplate3, "1999-01-05", config.Local)
	user := &User{
		Name:        "刘八",
		Description: "加个描述",
		Department: &Department{
			Name: "质量部",
		},
		Profile: Profile{
			Gender:   2,
			Birthday: birthday,
			Phones: []*Phone{
				{Value: "13456789012"},
				{Value: "13987654321"},
			},
		},
		Projects: []*Project{
			{Name: "检测手机"},
			{Name: "检测电脑"},
		},
		Roles: []*Role{
			{Name: "检测员"},
			{Name: "质量部主管"},
		},
	}
	err := user.Create(db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", user)
}

func TestUser_ModifyRole(t *testing.T) {
	user := &User{
		Id: 1,
	}
	newRoles := []*Role{
		{Id: 8},
		{Id: 5},
	}
	err := user.ModifyRole(db, newRoles)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}

func TestUser_Delete(t *testing.T) {
	user := &User{Id: 7}
	err := user.Delete(db)
	if err != nil {
		panic(err)
	}
}

func TestUser_Destroy(t *testing.T) {
	user := &User{Id: 7}
	err := user.Destroy(db)
	if err != nil {
		panic(err)
	}
}
