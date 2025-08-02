package auth

import (
	"context"

	uContr "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/contract/authUser"
	uAgg "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/domain/user_aggregate"
)



func fromCreateUserReq(ctx context.Context, cur *uContr.CreateUserRequest, id string) (*uAgg.User, error) {
	return uAgg.NewUser(
		ctx,
		id, 
		cur.Email,
		cur.Username,
		cur.Password,
	)
}

func toSignInResponse(token string, user *uAgg.User) *uContr.AuthResponse {
	return &uContr.AuthResponse{
		Token: token,
		User: uContr.SignInResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}
}