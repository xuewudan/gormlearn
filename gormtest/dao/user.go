package dao

import "log"

// User 定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
type User struct {
	ID int64 // 主键
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	//例如 `gorm:"column:username"` 标签说明含义是: Mysql表的列名（字段名)为username
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	//创建时间，时间戳
	CreateTime int64 `gorm:"column:createtime"`
}

func (u User) TableName() string {
	return "users"
}

func SaveUser(user *User) {
	err := DB.Create(user).Error
	if err != nil {
		log.Println("database connect err: ", err)
	}
}

func GetUserById(id int64) User {
	var user User
	err := DB.Where("id=?", id).First(&user).Error

	if err != nil {
		log.Println("find by id err: ", err)
	}
	return user
}

func GetAll() []User {
	var user []User
	err := DB.Find(&user).Error

	if err != nil {
		log.Println("find by id err: ", err)
	}
	return user
}

func UpdateUser(id int64) {
	err := DB.Model(&User{}).Where("id=?", id).Update("username", "dandan")
	if err != nil {
		log.Println("UpdateUser by id err: ", err)
	}
}

func DeleteUser(id int64) {
	err := DB.Where("id=?", id).Delete(&User{}).Error
	if err != nil {
		log.Println("UpdateUser by id err: ", err)
	}
}
