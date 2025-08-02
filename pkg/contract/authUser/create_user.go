package authuser

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender  string `json:"gender"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	// MetaData  MetaDataRes `json:"metadata"`
}

// type MetaDataRes struct {
// 	CreatedAt int64 `json:"created_at"`
// 	UpdatedAt int64 `json:"updated_at"`
// }

type CreateUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}


type SignInRequest struct {
	Username    string `json:"username" binding:"required"`
	Password 	string `json:"password" binding:"required"`
}

type SignInResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  SignInResponse `json:"user"`
}