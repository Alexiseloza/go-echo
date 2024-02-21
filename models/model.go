package models

type StatusModel struct {
	Status   string `json:"status"`
	Messagge string `json:"messagge"`
}

type CategoryModel struct {
	Name string `json:"name" validate:"required,min=2,max=25"`
	ID   int    `json:"id"`
}

type ProductModel struct {
	Name        string   `json:"name" validate:"required,min=3,max=20"`
	Price       int      `json:"price" validate:"required,numeric"`
	CategoryId  string   `json:"category_id" validate:"required"`
	Description string   `json:"description"`
	Image       []string `json:"image"`
	IsPromo     bool     `json:"is_promo"`
	Stock       int      `json:"stock" validate:"required,numeric"`
}

type UserModel struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type TokenResponse struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}
