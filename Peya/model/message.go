package model

//Message struct holds the data to be used with the user comments
type Message struct {
	UserID      int     `json:"userId"`
	OrderID     int     `json:"orderId"`
	ShopID      int     `json:"shopId"`
	Punctuation float64 `json:"punctuation"`
	Date        int     `json:"date"`
	Opinion     string  `json:"opinion"`
}
