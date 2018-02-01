package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	kmodel "github.com/relax-space/go-kit/model"
)

func Fruit_Get(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult("router param is missing or format is not correct."))
	}
	httpStatus, result, err := Fruit{}.Get(c.Request().Context(), id)
	if err != nil {
		return c.JSON(httpStatus, ErrorResult(err.Error()))
	}
	return c.JSON(http.StatusOK, SuccessResult(result))
}

func Fruit_Find(c echo.Context) error {
	var param struct {
		Limit int `query:"limit"`
		Start int `query:"start"`
	}
	if err := c.Bind(&param); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult("param error"))
	}
	if param.Limit == 0 {
		param.Limit = 30
	}
	status, totalCount, result, err := Fruit{}.Find(c.Request().Context(), param.Limit, param.Start)
	if err != nil {
		return c.JSON(status, ErrorResult(err.Error()))
	}
	return c.JSON(status, SuccessResult(kmodel.ArrayResult{TotalCount: totalCount, Items: result}))
}

func Fruit_Create(c echo.Context) error {
	fruits := new([]Fruit)
	if err := c.Bind(fruits); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}
	status, err := Fruit{}.CreateBatch(c.Request().Context(), fruits)
	if err != nil {
		return c.JSON(status, ErrorResult(err.Error()))
	}
	return c.JSON(status, SuccessResult(nil))
}

func Fruit_Update(c echo.Context) error {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}

	fruit := new(Fruit)
	if err := c.Bind(fruit); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult(err.Error()))
	}

	status, err := fruit.Update(c.Request().Context(), id)
	if err != nil {
		return c.JSON(status, ErrorResult(err.Error()))
	}
	return c.JSON(status, SuccessResult(nil))
}

func Fruit_Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResult("router param format is not correct."))
	}
	status, err := Fruit{}.Delete(c.Request().Context(), id)
	if err != nil {
		return c.JSON(status, ErrorResult(err.Error()))
	}
	return c.JSON(status, SuccessResult(nil))
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
