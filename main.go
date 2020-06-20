package main

import "GenericModel/source"

func main() {
	db := source.NewDB()
	//删除表
	dropTables(db)
	//创建表
	crateTables(db)
	//创建外键
	createForeignKey(db)
	//插入演示数据
	demoData(db)
}
