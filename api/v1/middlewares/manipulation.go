package middlewares

import (
	"github.com/NikitaNasevich/test_task_avito/api/v1/response"
	"github.com/gin-gonic/gin"
)

func GetBalanceMeddleware(c *gin.Context) {
	var err error
	var req struct {
		UserId int64 `binding:"required"`
	}
	if err = c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, response.NewErrorResponse(err.Error(), 1))
		return
	}

	if req.UserId <= 0 {
		c.AbortWithStatusJSON(400, response.NewErrorResponse("Bad user supplied", 2))
		return
	}

	c.Set("UserId", req.UserId)
}

func AddBalanceMeddleware(c *gin.Context) {
	var err error
	var req struct {
		UserId  int64   `binding:"required"`
		Balance float64 `binding:"required"`
	}
	if err = c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, response.NewErrorResponse(err.Error(), 1))
		return
	}

	if req.UserId <= 0 {
		c.AbortWithStatusJSON(400, response.NewErrorResponse("Bad user supplied", 2))
		return
	}

	c.Set("UserId", req.UserId)
	c.Set("Balance", req.Balance)
}

func ReserveFundsMeddleware(c *gin.Context) {
	var err error
	var req struct {
		UserId         int64 `binding:"required"`
		ServiceId      int64 `binding:"required"`
		OrderServiceId int64 `binding:"required"`
		ReserveBalance float64
	}
	if err = c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, response.NewErrorResponse(err.Error(), 1))
		return
	}

	if req.UserId <= 0 {
		c.AbortWithStatusJSON(400, response.NewErrorResponse("Bad user supplied", 2))
		return
	}

	if req.ServiceId <= 0 {
		c.AbortWithStatusJSON(400, response.NewErrorResponse("Bad service supplied", 3))
		return
	}

	if req.OrderServiceId <= 0 {
		c.AbortWithStatusJSON(400, response.NewErrorResponse("Bad order supplied", 4))
		return
	}

	c.Set("UserId", req.UserId)
	c.Set("ServiceId", req.ServiceId)
	c.Set("OrderServiceId", req.OrderServiceId)
	c.Set("ReserveBalance", req.ReserveBalance)

}

func AcceptProfitMeddleware(c *gin.Context) {
	var err error
	var req struct {
		UserId         int64 `binding:"required"`
		ServiceId      int64 `binding:"required"`
		OrderServiceId int64 `binding:"required"`
		ReserveBalance float64
	}
	if err = c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, response.NewErrorResponse(err.Error(), 1))
		return
	}

	if req.UserId <= 0 {
		c.AbortWithStatusJSON(400, response.NewErrorResponse("Bad user supplied", 2))
		return
	}

	if req.ServiceId <= 0 {
		c.AbortWithStatusJSON(400, response.NewErrorResponse("Bad service supplied", 3))
		return
	}

	if req.OrderServiceId <= 0 {
		c.AbortWithStatusJSON(400, response.NewErrorResponse("Bad order supplied", 4))
		return
	}

	c.Set("UserId", req.UserId)
	c.Set("ServiceId", req.ServiceId)
	c.Set("OrderServiceId", req.OrderServiceId)
	c.Set("Balance", req.ReserveBalance)

}
