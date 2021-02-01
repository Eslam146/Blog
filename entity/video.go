package entity


type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName string `json:"lastname" binding:"required"`
	Age string `json:"age" binding:"required"`
	Email string `json:"email" validate:"required,email"`
}
type Video struct{
	Title string  `json:"title" binding:"min=2,max=10"`
	Description string `json:"description" binding:"max=10"`
	URL string `json:"url" binding:"required,url"`
	Author string `json:"author" binding:"required"`
}