package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shudipta/coding-test/model"
	"github.com/shudipta/coding-test/repository"
)

func Hello(e echo.Context) error {
	return e.JSON(http.StatusOK, "Hello")
}

func SetRating(e echo.Context) error {
	orderdetail := &model.OrderDetail{}

	err := e.Bind(orderdetail)
	if err != nil {
		return e.JSON(http.StatusBadRequest, repository.ErrBadParams)
	}

	err = repository.UpdateOrderDetail(orderdetail)
	if err == repository.ErrNotFound {
		return e.JSON(http.StatusNotFound, err.Error())
	} else if err == repository.ErrUpdateFailed {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, "saved successfully")
}

func FetchTopCategories(e echo.Context) error {
	categories, err := repository.GetAllTopCategories()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, categories)
}
