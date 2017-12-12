package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/relax-space/go-kit/test"
)

func Test_Fruit_Create(t *testing.T) {
	f := &Fruit{
		Code: time.Now().Format("0102150405"),
		Name: "test",
	}
	status, err := f.Create(ctx)
	test.Ok(t, err)
	test.Equals(t, http.StatusCreated, status)
}
func Test_Fruit_Find(t *testing.T) {
	status, totalCount, result, err := Fruit{}.Find(ctx, 10, 0)
	fmt.Printf("\ntotalCount:%v,result:%+v", totalCount, result)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, status)
}

func Test_Fruit_Get(t *testing.T) {
	status, result, err := Fruit{}.Get(ctx, 12)
	fmt.Printf("\n%+v", result)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, status)

}

func Test_Fruit_Update(t *testing.T) {
	f := &Fruit{
		Id: 14,
		//Code: time.Now().Format("0102150405"),
		Name: "test update4",
	}

	status, err := f.Update(ctx)
	test.Ok(t, err)
	test.Equals(t, http.StatusNoContent, status)
}

func Test_Fruit_CreateBatch(t *testing.T) {
	f := Fruit{
		Code: time.Now().Format("0102150405"),
		Name: "test batch",
	}

	status, err := Fruit{}.CreateBatch(ctx, &[]Fruit{f})
	test.Ok(t, err)
	test.Equals(t, http.StatusCreated, status)
}

func Test_Fruit_Delete(t *testing.T) {
	status, err := Fruit{}.Delete(ctx, 13)
	test.Ok(t, err)
	test.Equals(t, http.StatusNoContent, status)
}
