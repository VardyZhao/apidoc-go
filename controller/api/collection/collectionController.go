package collection

import (
	"ginorm/controller"
	"ginorm/entity/request"
	"ginorm/service"
	"github.com/gin-gonic/gin"
)

func LoadPostman(c *gin.Context) {
	var req request.LoadPostmanRequest
	if err := c.ShouldBind(&req); err != nil {
		controller.FailWithErr(c, err)
		return
	}
	if err := req.Valid(); err != nil {
		controller.FailWithErr(c, err)
		return
	}

	serv := service.CollectionService{}
	if err := serv.LoadPostman(req); err != nil {
		controller.FailWithErr(c, err)
		return
	}

	res, err := serv.GetList(req.CollectionId)
	if err != nil {
		controller.FailWithErr(c, err)
		return
	}

	controller.Success(c, res)
}
