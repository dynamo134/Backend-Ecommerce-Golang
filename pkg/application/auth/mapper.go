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
		cur.Gender,
		cur.Phone,
		cur.Address,
	)
}