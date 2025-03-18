package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"go-quickstart/constant/request"
	"go-quickstart/domain"
	"net/http"
)

type TestController struct {
	testUsecase domain.TestUsecase
}

func NewTestController(tu domain.TestUsecase) *TestController {
	return &TestController{testUsecase: tu}
}

func (t *TestController) HelloPost(c *gin.Context) {
	c.JSON(http.StatusOK, &domain.SuccessResponse{Message: request.DefaultSuccessMessage, Code: request.DefaultSuccessCode})
}

func (t *TestController) HelloGet(c *gin.Context) {
	getParam := c.Param("get")
	if funk.IsEmpty(getParam) {
		c.JSON(http.StatusInternalServerError, &domain.ErrorResponse{Message: request.BindErrorMessage, Code: request.BindErrorCode})
	}
	ok := t.testUsecase.TestService()
	if ok {
		c.JSON(http.StatusOK, &domain.SuccessResponse{Message: request.DefaultSuccessMessage, Code: request.DefaultSuccessCode})
	} else {
		c.JSON(http.StatusInternalServerError, &domain.ErrorResponse{Message: request.DefaultErrorMessage, Code: request.DefaultErrorCode})
	}
}
