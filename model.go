package main

import (
	"context"
	"errors"
	"time"

	"github.com/go-xorm/xorm"

	"github.com/relax-space/go-kitt/factory"
)

type Fruit struct {
	Id        int64     `json:"id" xml:"id" form:"id" xorm:"int64 notnull autoincr pk 'id'"`
	CreatedAt time.Time `json:"created_at" xml:"created_at" form:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at" form:"updated_at" xorm:"updated"`
	Code      string    `json:"code" xml:"code" form:"code" xorm:"varchar(10)"`
	Name      string    `json:"name" xml:"name" form:"name" xorm:"nvarchar(50)"`
	Color     string    `json:"color" xml:"color" form:"color" xorm:"varchar(10)"`
	Price     int64     `json:"price" xml:"price" form:"price"`
	StoreCode string    `json:"store_code" xml:"store_code" form:"store_code" xorm:"varchar(10)"`
}

func (Fruit) TableName() string {
	return "fruit"
}

func (Fruit) Get(ctx context.Context, id int64) (fruit *Fruit, err error) {
	fruit = &Fruit{}
	has, err := factory.DB(ctx).ID(id).Get(fruit)
	if err != nil {
		return
	} else if has == false {
		err = errors.New("no data has found. ")
		return
	}
	return
}

func (Fruit) FindNames(ctx context.Context) (names []string, err error) {
	err = factory.DB(ctx).Table("fruit").Select("name").Find(&names)
	return
}

func (Fruit) Find(ctx context.Context, limit, start int) (totalCount int64, fruits []Fruit, err error) {
	session := *(factory.DB(ctx))
	errc := make(chan error)
	go func(session xorm.Session) {
		if totalCount, err = session.Count(&Fruit{}); err != nil {
			errc <- err
			return
		}
		errc <- nil
	}(session)
	go func(session xorm.Session) {
		if err = session.Limit(limit, start).Find(&fruits); err != nil {
			errc <- err
			return
		}
		errc <- nil
	}(session)

	if err = <-errc; err != nil {
		return
	}
	if err = <-errc; err != nil {
		return
	}
	return
}

func (f *Fruit) Update(ctx context.Context, id int64) (err error) {
	rowCount, err := factory.DB(ctx).Table("fruit").ID(id).Update(f)
	if err != nil {
		return
	} else if rowCount == 0 {
		err = errors.New("no data has changed.")
		return
	}
	return
}

func (f *Fruit) Create(ctx context.Context) (err error) {
	rowCount, err := factory.DB(ctx).InsertOne(f)
	if err != nil {
		return
	} else if rowCount == 0 {
		err = errors.New("no data has changed.")
		return
	}
	return
}

func (Fruit) CreateBatch(ctx context.Context, fruits *[]Fruit) (err error) {
	rowCount, err := factory.DB(ctx).Insert(fruits)
	if err != nil {
		return
	} else if rowCount == 0 {
		err = errors.New("no data has changed.")
		return
	}
	return
}

func (Fruit) Delete(ctx context.Context, id int64) (err error) {
	rowCount, err := factory.DB(ctx).ID(id).Delete(&Fruit{})
	if err != nil {
		return
	} else if rowCount == 0 {
		err = errors.New("no data has changed.")
		return
	}
	return
}

func setSortOrder(q *xorm.Session, sortby, order []string) error {
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				if order[i] == "desc" {
					q.Desc(v)
				} else if order[i] == "asc" {
					q.Asc(v)
				} else {
					return errors.New("Invalid order. Must be either [asc|desc]")
				}
			}
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				if order[0] == "desc" {
					q.Desc(v)
				} else if order[0] == "asc" {
					q.Asc(v)
				} else {
					return errors.New("Invalid order. Must be either [asc|desc]")
				}
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return errors.New("'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return errors.New("unused 'order' fields")
		}
	}
	return nil
}
