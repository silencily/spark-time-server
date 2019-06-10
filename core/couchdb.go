package core

import (
	"github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
	"golang.org/x/net/context"
)

var template *kivik.DB

func init() {
	client, err := kivik.New("couch", "http://localhost:5984/")
	if err != nil {
		panic(err)
	}
	err = client.Authenticate(context.TODO(), couchdb.BasicAuth("sparktime", "sparktime"))
	if err != nil {
		panic(err)
	}

	template = client.DB(context.TODO(), "sparktime")
	if template.Err() != nil {
		panic(err)
	}
}

//获取操作数据库的模板对象
func GetCouchDBTemplate() *kivik.DB {
	return template
}
