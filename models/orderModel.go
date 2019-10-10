package models

//Order db model
type Order struct {
	ID            int     `db:"id"`
	OrderStatusID int     `db:"orderStatusId"`
	OrderMessage  string  `db:"orderMessage"`
	OrderAmount   float64 `db:"orderAmount"`
}

//OrderStatus db model
type OrderStatus struct {
	ID              int    `db:"id"`
	OrderStatusName string `db:"orderStatusNameo"`
	Order           []Order
}

//OrderResponseModel ...
type OrderResponseModel struct {
	ID              int     `json:"id"`
	OrderStatusName string  `json:"orderStatusName"`
	OrderMessage    string  `json:"orderMessage"`
	OrderAmount     float64 `json:"orderAmount"`
}
