package main

import (
	"context"
	"os"

	"github.com/go-xorm/xorm"
	"github.com/relax-space/go-kitt/echomiddleware"
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
	ctx = context.WithValue(context.Background(), echomiddleware.ContextDBName, xormEngine.NewSession())
}
