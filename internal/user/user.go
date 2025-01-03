package user

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name" binding:"required,max=255"`
	Email string `json:"email" binding:"required,email,max=255"`
}
