package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/relax-space/go-kit/test"
)

func Test_Fruit_Create(t *testing.T) {
	f := &Fruit{
		Code: time.Now().Format("0102150405"),
		Name: "test",
	}
	err := f.Create(ctx)
	test.Ok(t, err)
}
func Test_Fruit_Find(t *testing.T) {
	totalCount, result, err := Fruit{}.Find(ctx, 10, 0)
	fmt.Printf("\ntotalCount:%v,result:%+v", totalCount, result)
	test.Ok(t, err)
}

func Test_Fruit_Get(t *testing.T) {
	result, err := Fruit{}.Get(ctx, 12)
	fmt.Printf("\n%+v", result)
	test.Ok(t, err)

}

func Test_Fruit_GetName(t *testing.T) {
	result, err := Fruit{}.FindNames(ctx)
	fmt.Printf("\n%+v", result)
	test.Ok(t, err)

}

func Test_Fruit_Update(t *testing.T) {
	f := &Fruit{
		Name: "test update4",
	}

	err := f.Update(ctx, 15)
	test.Ok(t, err)
}

func Test_Fruit_CreateBatch(t *testing.T) {
	f := Fruit{
		Code: time.Now().Format("0102150405"),
		Name: "test batch",
	}

	err := Fruit{}.CreateBatch(ctx, &[]Fruit{f})
	test.Ok(t, err)
}

func Test_Fruit_Delete(t *testing.T) {
	err := Fruit{}.Delete(ctx, 18)
	test.Ok(t, err)
}
