package auth

import (
	"context"
	"fmt"

	uContr "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/contract/authUser"
	aRepo "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/infrastructure/persistence/auth"
	jwtUtils "github.com/dynamo134/Backend-Ecommerce-Golang/pkg/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepo *aRepo.AuthRepository
}

// NewAuthService creates a new instance of AuthService with the provided AuthRepository.
func NewAuthService(
	aRepo *aRepo.AuthRepository,

	) *AuthService {
	return &AuthService{
		authRepo: aRepo,
	}
}

func (s *AuthService) SignUp(ctx context.Context, cur *uContr.CreateUserRequest) (*uContr.CreateUserResponse, error) {

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cur.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	cur.Password = string(hashedPassword)

	id := uuid.New().String()


	dbUser, err := fromCreateUserReq(ctx, cur, id)
	err = s.authRepo.CreateUser(ctx, dbUser)
	if err != nil {
		return nil, err
	}

	return &uContr.CreateUserResponse{
		ID:    dbUser.ID,
		Name:  dbUser.Name,
		Email: dbUser.Email,
	}, nil
}


func (s *AuthService) SignIn(ctx context.Context, sr *uContr.SignInRequest) (*uContr.AuthResponse, error) {
	// Fetch user from the repository
	user, err := s.authRepo.GetUserByUsername(ctx, sr.Username)
	if err != nil {
		return nil, err
	}
	
	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(sr.Password))
	if err != nil {
		return nil, err
	}
	// Generate JWT token
	token, err := jwtUtils.CreateJWT(user.Name, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create JWT token: %w", err)
	}
	fmt.Printf("✅ User %s signed in successfully\n", user.Name)
	// Return the user information
	return toSignInResponse(token, user), nil
}

