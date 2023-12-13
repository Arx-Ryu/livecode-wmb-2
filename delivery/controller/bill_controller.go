package controller

import (
	"livecode-wmb-rest-api/delivery/api"
	"livecode-wmb-rest-api/delivery/middleware"
	"livecode-wmb-rest-api/manager"
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BillController struct {
	router 				*gin.Engine
	ucBill				manager.UseCaseManager
	api.BaseApi
}

func (con *BillController) createBill(c *gin.Context) {
	var item *model.Bill
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	err = con.ucBill.CreateBillUseCase().CreateBill(item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *BillController) printBill(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		con.Failed(c, err)
		return
	}
	bill, err := con.ucBill.PrintBillUseCase().PrintBill(id)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, bill)
}

func (con *BillController) createTable(c *gin.Context) {
	var item *model.Table
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	err = con.ucBill.TableCRUDUseCase().Create(item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *BillController) readTable(c *gin.Context) {
	item, err := con.ucBill.TableCRUDUseCase().Read()
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *BillController) updateTable(c *gin.Context) {
	var item *model.Table
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	oldItem := model.Table{
		Id: item.Id,
	}
	log.Print(item)
	err = con.ucBill.TableCRUDUseCase().Update(&oldItem, *item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *BillController) deleteTable(c *gin.Context) {
	var item *model.Table
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	err = con.ucBill.TableCRUDUseCase().Delete(item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *BillController) createTransType(c *gin.Context) {
	var item *model.TransType
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	err = con.ucBill.TransTypeCRUDUseCase().Create(item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *BillController) readTransType(c *gin.Context) {
	item, err := con.ucBill.TransTypeCRUDUseCase().Read()
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *BillController) updateTransType(c *gin.Context) {
	var item *model.TransType
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	oldItem := model.TransType{
		Id: item.Id,
	}
	err = con.ucBill.TransTypeCRUDUseCase().Update(&oldItem, *item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *BillController) deleteTransType(c *gin.Context) {
	var item *model.TransType
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	err = con.ucBill.TransTypeCRUDUseCase().Delete(item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *BillController) dailySales(c *gin.Context) {
	item, err := con.ucBill.DailySalesUseCase().DailySales()
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, item)
}

func (con *BillController) payment(c *gin.Context) {
	var item *model.DtoPaymentLopei
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	status, err := con.ucBill.LopeiUseCase().DoPayment(int32(item.LopeiId), item.Amount)	
	if !status {
		con.Failed(c, err)
		return
	}
	if err != nil {
		con.Failed(c, err)
		return
	}
	bill, err := con.ucBill.PrintBillUseCase().PrintBill(item.BillId)
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, bill)
}

func (con *BillController) balance(c *gin.Context) {
	var item *model.DtoPaymentLopei
	err := con.ParseRequestBody(c, &item)
	if err != nil {
		con.Failed(c, err)
		return
	}
	balance, err := con.ucBill.LopeiUseCase().GetBalance(int32(item.LopeiId))
	if err != nil {
		con.Failed(c, err)
		return
	}
	con.Success(c, balance)
}

func (con *BillController) login(c *gin.Context) {
	var user model.Credential
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Cant Bind Struct",
		})
		return 
	}
	token, err := con.ucBill.LoginUseCase().Login(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":err,
		})
		return
	}
	if token.AccessToken != "" {
		c.JSON(http.StatusOK, gin.H{
			"token":token,
		})
		return
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}


func NewBillController(router *gin.Engine, ucBill manager.UseCaseManager, t utils.Token) *BillController {
	controller := BillController{
		router: router,
		ucBill: ucBill,
	}

	router.POST("/table", controller.createTable)
	router.GET("/table", controller.readTable)
	router.PUT("/table", controller.updateTable)
	router.DELETE("/table", controller.deleteTable)

	router.POST("/transtype", controller.createTransType)
	router.GET("/transtype", controller.readTransType)
	router.PUT("/transtype", controller.updateTransType)
	router.DELETE("/transtype", controller.deleteTransType)

	router.POST("/login", controller.login)
	
	protected := router.Group("/cashier", middleware.NewTokenValidator(t).RequireToken())

	protected.GET("/balanceLopei", controller.balance)
	protected.GET("/paymentLopei", controller.payment)
	protected.POST("/bill", controller.createBill)
	protected.GET("/bill/:id", controller.printBill)
	protected.GET("/sales", controller.dailySales)

	return &controller
}