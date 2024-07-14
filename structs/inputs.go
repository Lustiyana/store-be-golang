package structs

type LoginInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	FullName string `json:"fullname" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Address string `json:"address"`
	Province string `json:"province"`
	City string `json:"city"`
	PostalCode string `json:"postal_code"`
	Country string `json:"country"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type ProductInput struct {
	CategoryID uint `json:"category_id"`
	Description string `json:"description"`
	Price uint `json:"price"`
	UserID uint `json:"user_id"`
	ProductName string `json: "product_name"`
}

type ImageInput struct {
	ProductID uint `json:"product_id"`
	Alt string `json:"alt"`
	Url string `json:"url"`
}

type OrderInput struct {
	UserID uint `json:"user_id" binding:"required"`
	Address string `json:"address" binding:"required"`
	ProductID uint `json:"product_id" binding:"required"`
}

type UpdatePaymentInput struct {
	OrderID uint `json:"order_id" binding:"required"`
	PaymentStatus string `json:"payment_status" binding:"required"` 
}