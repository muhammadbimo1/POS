package delivery

import (
	"table_management/delivery/appresponse"
	"table_management/dto"
	"table_management/usecase"

	"github.com/gin-gonic/gin"
)

type OrderApi struct {
	usecase     usecase.OrderUseCase
	publicRoute *gin.RouterGroup
}

func NewOrderApi(publicRoute *gin.RouterGroup, useCase usecase.OrderUseCase) *OrderApi {
	api := OrderApi{
		usecase:     useCase,
		publicRoute: publicRoute,
	}
	api.initRouter()
	return &api
}

func (o *OrderApi) initRouter() {
	tableRoute := o.publicRoute.Group("/order")
	tableRoute.POST("/open", o.openBill)
	tableRoute.POST("/close", o.closeBill)

}

func (o *OrderApi) openBill(c *gin.Context) {
	var orderRequest dto.OrderRequest
	response := appresponse.NewJsonResponse(c)
	err := c.BindJSON(&orderRequest)
	if err != nil {
		response.SendError(*appresponse.NewBadRequestError(err, "bind json error"))
		return
	}
	order, err := o.usecase.OpenBill(orderRequest)
	if err != nil {
		response.SendError(*appresponse.NewInternalServerError(err, "Something went in openbill usecase"))
		return
	}
	response.SendData(&appresponse.ResponseMessage{Status: "SUCCESS", Description: "order", Data: order})
}

func (o *OrderApi) closeBill(c *gin.Context) {
	var closeBillRequest dto.CloseBillRequest
	response := appresponse.NewJsonResponse(c)
	err := c.BindJSON(&closeBillRequest)
	if err != nil {
		response.SendError(*appresponse.NewBadRequestError(err, "bind json error"))
		return
	}
	closeBillRequest.Total, err = o.usecase.GetTotal(closeBillRequest.BillNo)
	if err != nil {
		response.SendError(*appresponse.NewInternalServerError(err, "Something went in get total"))
		return
	}
	order, err := o.usecase.CloseBill(closeBillRequest)
	if err != nil {
		response.SendError(*appresponse.NewInternalServerError(err, "Something went in closebill  usecase"))
		return
	}
	response.SendData(&appresponse.ResponseMessage{Status: "SUCCESS", Description: "closebill", Data: order})
}
