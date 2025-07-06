package authuser

type CreateUserRequest struct {
	ID 	 string `json:"id"`
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