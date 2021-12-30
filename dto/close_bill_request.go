package dto

type CloseBillRequest struct {
	BillNo        string
	PaymentMethod string
	PhoneNo       string
	Total         int
}
