package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User 用户 模型
type User struct {
	DbBase `bson:",omitempty"`
	ID     bson.ObjectId `bson:"_id,omitempty" json:"id"`

	Username    string    `bson:"username,omitempty" json:"username"`       // Username 用户名
	Password    string    `bson:"password,omitempty" json:"password"`       // Password 摘要后的密码
	Email       string    `bson:"email,omitempty" json:"email"`             // Email 邮箱
	IsAdmin     bool      `bson:"isAdmin,omitempty" json:"isAdmin"`         // IsAdmin 是否为管理员
	IsAbleLogin bool      `bson:"isAbleLogin,omitempty" json:"isAbleLogin"` // IsAbleLogin 是否可以登录
	CreateTime  time.Time `bson:"createTime,omitempty" json:"createTime"`   // CreateTime 创建时间
}

// CName 获取当前集合名称
func (a *User) CName() string {
	return "user"
}

// Insert 插入记录
func (a *User) Insert() error {
	return a.Collection(a.CName()).Insert(a)
}

// DeleteByID 删除记录 根据id
func (a *User) DeleteByID() error {
	return a.Collection(a.CName()).RemoveId(a.ID)
}

// UpdateByID 更新记录 根据id
func (a *User) UpdateByID() error {
	return a.Collection(a.CName()).UpdateId(a.ID, a)
}

// GetOneByID 根据id 获取单条数据
func (a *User) GetOneByID() (User, error) {
	var user User
	err := a.Collection(a.CName()).FindId(a.ID).One(&user)
	return user, err
}

// GetOneByUsername 根据id 获取单条数据
func (a *User) GetOneByUsername() (User, error) {
	var user User
	err := a.Collection(a.CName()).Find(bson.M{"username": a.Username}).One(&user)
	return user, err
}

// GetOneByEmail 根据id 获取单条数据
func (a *User) GetOneByEmail() (User, error) {
	var user User
	err := a.Collection(a.CName()).Find(bson.M{"email": a.Email}).One(&user)
	return user, err
}

// GetCountByUsernameOrEmail 根据 用户名 或 邮箱 获取记录条数
func (a *User) GetCountByUsernameOrEmail() (int, error) {
	return a.Collection(a.CName()).
		Find(bson.M{"$or": []bson.M{bson.M{"username": a.Username}, bson.M{"email": a.Email}}}).
		Count()
}

// GetTotalCount 总用户数
func (a *User) GetTotalCount() (int, error) {
	return a.Collection(a.CName()).Count()
}
