package main

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/relax-space/go-kitt/echomiddleware"
)

type Fruit struct {
	Id        int64     `json:"id" xorm:"int64 notnull autoincr pk 'id'"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
	Code      string    `json:"code" xorm:"varchar(10)"`
	Name      string    `json:"name" xorm:"nvarchar(50)"`
	Color     string    `json:"color" xorm:"varchar(10)"`
	Price     int64     `json:"price"`
	StoreCode string    `json:"store_code" xorm:"varchar(10)"`
}

func (Fruit) TableName() string {
	return "fruit"
}

func (Fruit) Get(ctx context.Context, id int64) (httpStatus int, fruit *Fruit, err error) {
	httpStatus = http.StatusOK
	fruit = &Fruit{}
	has, err := echomiddleware.DB(ctx).ID(id).Get(fruit)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	} else if has == false {
		err = errors.New("no data has found.")
		return
	}
	return
}
