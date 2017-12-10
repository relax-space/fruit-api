package main

import (
	"fmt"
	"kit/test"
	"net/http"
	"testing"
)

func Test_Fruit_Get(t *testing.T) {
	status, result, err := Fruit{}.Get(ctx, 2)
	fmt.Printf("\n%+v", result)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, status)

}
