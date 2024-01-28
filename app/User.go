package app

type RegisterInput struct {
    Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateInput struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}