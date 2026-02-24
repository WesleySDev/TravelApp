// Struct para receber dados (Request)
package controllers

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
