package model

import (
	"fmt"

	config "github.com/ClimbingMoon/camm/config"
	log "github.com/sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
)

var (
	// Session TODO 不应对包外暴露Session
	_Session  *mgo.Session
	_Database *mgo.Database
)

func init() {
	// TODO mongodb user password
	mongoURL := fmt.Sprintf("mongodb://%s:%s", config.ConfigSource.MongoHost, config.ConfigSource.MongoPort)
	log.WithField("mongo_url", mongoURL).Info("Connecting MongoDB...")
	Session, err := mgo.Dial(mongoURL)
	if err != nil {
		panic(fmt.Sprintf("Initialize mongodb error: %v", err))
	}
	_Database = Session.DB(config.ConfigSource.MongoDatabase)
	if err = Session.Ping(); err != nil {
		panic(fmt.Sprintf("MongoDB execute ping error: %v", err))
	}
	log.Info("MongoDB initialize success.")

	mgo.SetDebug(config.ConfigSource.Debug)
}

// DbBase 提供操作数据库的基类
type DbBase struct{}

// Session 获取当前操作数据库Session
func (d *DbBase) Session() *mgo.Session {
	return _Session
}

// Database 获取当前操作数据库
func (d *DbBase) Database() *mgo.Database {
	return _Database
}

// Collection 获取操作集合
func (d *DbBase) Collection(cName string) *mgo.Collection {
	return d.Database().C(cName)
}

// Find 提供简单的Find查询
func (d *DbBase) Find(cName string, query, selector interface{}) *mgo.Query {
	return d.Database().C(cName).Find(query).Select(selector)
}
