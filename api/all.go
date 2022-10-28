package api

import (
	v1 "github.com/NikitaNasevich/test_task_avito/api/v1"
	"github.com/gin-gonic/gin"
)

func AllApi(r *gin.Engine) {
	v1.GetBalanceApi(r)
	v1.AddBalanceApi(r)
	v1.ReserveFundsApi(r)
	v1.AcceptProfitApi(r)
}
