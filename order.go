package go_order_processing_queue

type OrderData struct {
	UserName   string  `json:"user_name"`
	OrderID    string  `json:"order_id"`
	OrderType  string  `json:"order_type"`
	OrderState int     `json:"order_state"` // 0.default 1.processing 2.complete 3.rollback
	Memo       string  `json:"memo"`
	Money      float64 `json:"money"`
}
