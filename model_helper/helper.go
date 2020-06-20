package model_helper

import "github.com/jinzhu/gorm"

/*
 * Description: 创建外键
 * Args:
 *  - db:
 *  - model:模型
 *  - field:字段
 *  - dest:关联表的字段
 *  - onDelete:删除时的事件
 *  - onUpdate:更新时的事件
 * Returns:
 *  - error:
 */
func CreateForeignKey(db *gorm.DB, model interface{}, field string, dest string, onDelete string, onUpdate string) error {
	return db.Model(model).AddForeignKey(field, dest, onDelete, onUpdate).Error
}
