package models

import "time"

type GetBalanceEntry struct {
	UserId  int64   `db:"customer_id"`
	Balance float64 `db:"balance"`
}

type AddBalanceEntry struct {
	UserId  int64
	Balance float64
}

type ReserveFundsEntry struct {
	UserId         int64
	ServiceId      int64
	OrderServiceId int64
	Balance        float64
}

type AcceptProfitEntry struct {
	UserId         int64
	ServiceId      int64
	OrderServiceId int64
	Date           time.Time
	Balance        float64
}
