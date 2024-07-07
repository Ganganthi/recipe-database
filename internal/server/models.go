package server

type userBase struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Username  string `json:"username" binding:"required"`
}

type UserCreate struct {
	userBase
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UserData struct {
	userBase
	ID string `json:"id" binding:"required"`
}
