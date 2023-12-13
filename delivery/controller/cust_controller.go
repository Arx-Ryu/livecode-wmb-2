package controller

import (
	"livecode-wmb-rest-api/delivery/api"
	"livecode-wmb-rest-api/manager"
	"livecode-wmb-rest-api/model"
	"strconv"
	"github.com/gin-gonic/gin"
)

type CustController struct {
	router 				*gin.Engine
	ucCust				manager.UseCaseManager
	api.BaseApi
}

func (con *CustController) registerCustomer(c *gin.Context) {
	var item *model.Customer
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	err = con.ucCust.CustomerRegistrationUseCase().Register(item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *CustController) findCustomer(c *gin.Context) {
	var item []model.Customer
	item, err := con.ucCust.CustomerRegistrationUseCase().GetCustomer()
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *CustController) activateMember(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		con.Failed(c, err)
		return
	}
	cust, err := con.ucCust.CustomerRegistrationUseCase().ActivateMember(id)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, cust)
}

func (con *CustController) createDiscount(c *gin.Context) {
	var item *model.Discount
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	err = con.ucCust.DiscCRUDUseCase().Create(item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *CustController) readDiscount(c *gin.Context) {
	item, err := con.ucCust.DiscCRUDUseCase().Read()
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *CustController) updateDiscount(c *gin.Context) {
	var item *model.Discount
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	oldItem := model.Discount{
		Id: item.Id,
	}
	err = con.ucCust.DiscCRUDUseCase().Update(&oldItem, *item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *CustController) deleteDiscount(c *gin.Context) {
	var item *model.Discount
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	err = con.ucCust.DiscCRUDUseCase().Delete(item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func NewCustController(router *gin.Engine, ucCust manager.UseCaseManager) *CustController {
	controller := CustController{
		router: router,
		ucCust: ucCust,
	}
	router.POST("/customer", controller.registerCustomer)
	router.GET("/customer", controller.findCustomer)
	router.PUT("/customer/:id", controller.activateMember)

	router.POST("/discount", controller.createDiscount)
	router.GET("/discount", controller.readDiscount)
	router.PUT("/discount", controller.updateDiscount)
	router.DELETE("/discount", controller.deleteDiscount)

	return &controller
}