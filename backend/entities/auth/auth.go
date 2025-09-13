package auth

type FormatRequest struct {
  Email    string `json:"email" binding:"required,email"`
  Password string `json:"password" binding:"required"`
}

type FormatResponse struct {
  Data interface{} `json:"data"`
}