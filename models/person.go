package models 

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Age       int    `json:"age" binding:"required,gte=0,lte=120"`
	Gender    string `json:"gender" binding:"required,oneof=male female other"`
}