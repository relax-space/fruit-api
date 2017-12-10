package main

import (
	"context"
	"os"

	"github.com/go-xorm/xorm"
)

var ctx context.Context

func init() {

	xormEngine, err := xorm.NewEngine("mysql", os.Getenv("FRUIT_CONN"))
	if err != nil {
		panic(err)
	}
	xormEngine.ShowSQL(true)
	//defer xormEngine.Close()
	xormEngine.Sync(new(Fruit))
	ctx = context.WithValue(context.Background(), ContextDBName, xormEngine.NewSession())
}
