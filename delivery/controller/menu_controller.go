package controller

import (
	"livecode-wmb-rest-api/delivery/api"
	"livecode-wmb-rest-api/manager"
	"livecode-wmb-rest-api/model"
	"log"
	"github.com/gin-gonic/gin"
)

type MenuController struct {
	router 				*gin.Engine
	ucMenu				manager.UseCaseManager
	api.BaseApi
}

func (con *MenuController) create(c *gin.Context) {
	var item *model.Menu
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	err = con.ucMenu.MenuCRUDUseCase().Create(item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *MenuController) read(c *gin.Context) {
	item, err := con.ucMenu.MenuCRUDUseCase().Read()
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *MenuController) update(c *gin.Context) {
	var item *model.Menu
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	oldItem := model.Menu{
		Id: item.Id,
	}
	err = con.ucMenu.MenuCRUDUseCase().Update(&oldItem, *item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *MenuController) delete(c *gin.Context) {
	var item *model.Menu
	err := con.ParseRequestBody(c, &item)
	log.Print(item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	err = con.ucMenu.MenuCRUDUseCase().Delete(item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *MenuController) createMPrice(c *gin.Context) {
	var item *model.MenuPrice
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	err = con.ucMenu.MPriceCRUDUseCase().Create(item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *MenuController) readMPrice(c *gin.Context) {
	item, err := con.ucMenu.MPriceCRUDUseCase().Read()
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *MenuController) updateMPrice(c *gin.Context) {
	var item *model.MenuPrice
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	oldItem := model.MenuPrice{
		Id: item.Id,
	}
	err = con.ucMenu.MPriceCRUDUseCase().Update(&oldItem, *item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *MenuController) deleteMPrice(c *gin.Context) {
	var item *model.MenuPrice
	err := con.ParseRequestBody(c, &item)
	log.Print(item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	err = con.ucMenu.MPriceCRUDUseCase().Delete(item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func NewMenuController(router *gin.Engine, ucMenu manager.UseCaseManager) *MenuController {
	controller := MenuController{
		router: router,
		ucMenu: ucMenu,
	}
	router.POST("/menu", controller.create)
	router.GET("/menu", controller.read)
	router.PUT("/menu", controller.update)
	router.DELETE("/menu", controller.delete)

	router.POST("/mprice", controller.createMPrice)
	router.GET("/mprice", controller.readMPrice)
	router.PUT("/mprice", controller.updateMPrice)
	router.DELETE("/mprice", controller.deleteMPrice)

	return &controller
}