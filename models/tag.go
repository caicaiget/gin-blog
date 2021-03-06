package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  int64 `json:"createdBy"`
	ModifiedBy int64 `json:"modifiedBy"`
	State      int64    `json:"state"`
}

func GetTags(pageNum int64, pageSize int64, userId int64) (tags []Tag) {
	db.Where("Created_by = ?", userId).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(userId int64) (count int64) {
	db.Model(&Tag{}).Where("Created_by = ?", userId).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int64, createBy int64) (tag *Tag, err error) {
	tag = &Tag{
		Name:      name,
		State:     state,
		CreatedBy: createBy,
		ModifiedBy: createBy,
	}
	err = db.Create(tag).Error
	return
}

func (tag Tag) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreatedOn", time.Now().Format("2006-01-02 15:04:05"))
	_ = scope.SetColumn("ModifiedOn", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ModifiedOn", time.Now().Format("2006-01-02 15:04:05"))

	return nil
}

func ExistTagByID(id int64) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func DeleteTag(id int64) bool {
	db.Exec("Update blog_article set is_deleted = 1 where id = ?", id)
	return true
}

func EditTag(id int64, data interface{}) (*Tag, error) {
	var tag Tag
	err := db.Model(&tag).Where("id = ?", id).Updates(data).Error
	log.Println(tag)
	return &tag, err
}
