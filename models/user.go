package models

import "github.com/jinzhu/gorm"

//用户表
type User struct {
	gorm.Model
	Name   string `gorm:"unique_index"`
	Email  string `gorm:"unique_index"`
	Avatar string
	Pwd    string
	Role   int `gorm:"default:1"` // 0 管理员 1正常用户
}

func (db *DB) QueryUserByEmailAndPassword(email, password string) (user User, err error) {
	return user, db.db.Model(&User{}).Where("email = ? and pwd = ?", email, password).Take(&user).Error
}

func (db *DB) QueryUserByName(name string) (user User, err error) {
	return user, db.db.Where("name = ?", name).Take(&user).Error
}

func (db *DB) QueryUserByEmail(email string) (user User, err error) {
	return user, db.db.Where("email = ?", email).Take(&user).Error
}

func (db *DB) UpdateUserEditor(editor string) (err error) {
	return db.db.Model(&User{}).Update("editor", editor).Error
}

func SaveUser(user *User) error {
	return db.Create(user).Error
}
