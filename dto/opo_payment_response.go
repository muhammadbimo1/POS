package dto

type OpoPaymentResponse struct {
	Message string
	Data    OpoReceipt
}

type OpoReceipt struct {
	ReceiptID string `json:"receipt_id"`
}
