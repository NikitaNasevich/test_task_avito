package v1

import (
	"database/sql"
	"github.com/NikitaNasevich/test_task_avito/api/v1/middlewares"
	"github.com/NikitaNasevich/test_task_avito/api/v1/models"
	"github.com/NikitaNasevich/test_task_avito/api/v1/response"
	"github.com/NikitaNasevich/test_task_avito/db"
	"github.com/gin-gonic/gin"
	"time"
)

func GetBalanceApi(r *gin.Engine) {
	group := r.Group("v1")

	group.Use(middlewares.GetBalanceMeddleware)

	group.POST("getBalance", func(c *gin.Context) {

		userId := c.GetInt64("UserId")

		var record models.GetBalanceEntry
		var err error
		if err = db.Database().Get(&record, "SELECT customer_id, balance \nFROM customer\nWHERE customer_id = ?", userId); err != nil && err != sql.ErrNoRows {
			c.AbortWithStatusJSON(500, response.NewErrorResponse(err.Error(), 100))
			return
		}

		if record.UserId == 0 {
			c.AbortWithStatusJSON(500, response.NewErrorResponse("UserId not found", 101))
			return
		}
		c.JSON(200, record)
	})

}

func AddBalanceApi(r *gin.Engine) {

	group := r.Group("v1")
	group.Use(middlewares.AddBalanceMeddleware)
	group.POST("addBalance", func(c *gin.Context) {

		userId := c.GetInt64("UserId")
		addBalance := c.GetFloat64("Balance")

		var record models.GetBalanceEntry
		var err error
		if err = db.Database().Get(&record, "SELECT customer_id, IF(balance IS NULL, 0, balance) AS balance\nFROM customer\nWHERE customer_id = ?", userId); err != nil && err != sql.ErrNoRows {
			c.AbortWithStatusJSON(500, response.NewErrorResponse(err.Error(), 102))
			return
		}

		var record2 models.AddBalanceEntry
		record2.Balance = record.Balance + addBalance
		record2.UserId = record.UserId

		if _, err = db.Database().NamedExec("UPDATE customer SET balance=:Balance WHERE customer_id= :UserId", record2); err != nil {
			c.AbortWithStatusJSON(500, response.NewErrorResponse(err.Error(), 103))
			return
		}

	})
}

func ReserveFundsApi(r *gin.Engine) {

	group := r.Group("v1")
	group.Use(middlewares.ReserveFundsMeddleware)
	group.POST("reserveFunds", func(c *gin.Context) {

		userId := c.GetInt64("UserId")
		serviceId := c.GetInt64("ServiceId")
		orderServiceId := c.GetInt64("OrderServiceId")
		reserveBalance := c.GetFloat64("ReserveBalance")

		var recordWithdraw models.AddBalanceEntry
		var err error

		recordWithdraw.UserId = userId
		recordWithdraw.Balance = reserveBalance

		if _, err = db.Database().NamedExec("UPDATE customer SET balance= balance - :Balance WHERE customer_id= :UserId", recordWithdraw); err != nil {
			c.AbortWithStatusJSON(500, response.NewErrorResponse(err.Error(), 104))
			return
		} else {
			var recordReserve models.ReserveFundsEntry
			recordReserve.UserId = userId
			recordReserve.ServiceId = serviceId
			recordReserve.OrderServiceId = orderServiceId
			recordReserve.Balance = reserveBalance

			if _, err = db.Database().NamedExec("INSERT INTO reserve (customer_id, service_id, order_id, summ)\nVALUES (:UserId, :ServiceId, :OrderServiceId, :Balance)\n    ", recordReserve); err != nil {
				c.AbortWithStatusJSON(500, response.NewErrorResponse(err.Error(), 105))
				return
			}
		}

	})
}

func AcceptProfitApi(r *gin.Engine) {

	group := r.Group("v1")
	group.Use(middlewares.AcceptProfitMeddleware)
	group.POST("acceptProfit", func(c *gin.Context) {
		var err error

		var acceptProfit models.AcceptProfitEntry
		acceptProfit.UserId = c.GetInt64("UserId")
		acceptProfit.ServiceId = c.GetInt64("ServiceId")
		acceptProfit.OrderServiceId = c.GetInt64("OrderServiceId")
		acceptProfit.Balance = c.GetFloat64("Balance")
		acceptProfit.Date = time.Now()

		if _, err = db.Database().NamedExec("INSERT INTO profit (receiving_date, summ, order_id)\nVALUES (:Date, :Balance, (SELECT order_id \n                         FROM reserve\n                         WHERE service_id = :ServiceId\n                         AND customer_id = :UserId\n                         AND summ = :Balance\n                         LIMIT 1));\n", acceptProfit); err != nil {
			c.AbortWithStatusJSON(500, response.NewErrorResponse(err.Error(), 106))
			return
		} else {
			if _, err = db.Database().NamedExec("UPDATE reserve \nSET summ = summ - :Balance\nWHERE customer_id = :UserId\nAND service_id = :ServiceId\nAND summ = :Balance\n", acceptProfit); err != nil {
				c.AbortWithStatusJSON(500, response.NewErrorResponse(err.Error(), 107))
				return
			}
		}

	})
}
