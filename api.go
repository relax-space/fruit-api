package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	kmodel "github.com/relax-space/go-kit/model"
)

func Fruit_Get(c echo.Context) error {
	var id int64
	if idStr := c.Param("id"); len(idStr) == 0 {
		fmt.Println(idStr)
		return c.JSON(http.StatusBadRequest, ErrorResult("router param is missing."))
	} else {
		var err error
		if id, err = strconv.ParseInt(idStr, 10, 64); err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResult("router param format is not correct."))
		}
	}
	httpStatus, result, err := Fruit{}.Get(c.Request().Context(), id)
	if err != nil {
		return c.JSON(httpStatus, ErrorResult(err.Error()))
	}
	return c.JSON(http.StatusOK, SuccessResult(result))
}

func ErrorResult(errMsg string) (result kmodel.Result) {
	result = kmodel.Result{
		Error: kmodel.Error{Message: errMsg},
	}
	return
}
func SuccessResult(param interface{}) (result kmodel.Result) {
	result = kmodel.Result{
		Success: true,
		Result:  param,
	}
	return
}
