package usecase

import (
	"cleanArch/todos/services/auth"
	"cleanArch/todos/services/model"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type AuthClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	UserId   string `json:"userId"`
}
type authUseCase struct {
	userRepo       auth.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo auth.UserRepository,
	hashSalt string,
	signingKey []byte,
	tokenTTL int64) auth.UseCase {
	return &authUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * time.Duration(tokenTTL),
	}

}
func (au *authUseCase) SignUp(ctx context.Context, username, password string, limit int) (*model.User, error) {
	fmtusername := strings.ToLower(username)
	euser, _ := au.userRepo.GetUserByUsername(ctx, fmtusername)
	if euser != nil {
		return nil, auth.ErruserExisted
	}
	user := &model.User{
		Id:       uuid.New().String(),
		Username: fmtusername,
		Password: password,
		Limit:    limit,
	}
	user.HashPassword()
	err := au.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return au.userRepo.GetUserByUsername(ctx, username)
}
func (au *authUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	user, _ := au.userRepo.GetUserByUsername(ctx, username)
	if user == nil {
		return "", auth.ErrUserNotFount
	}
	claims := AuthClaims{
		Username: user.Username,
		UserId:   user.Id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    "go-todos",
			ExpiresAt: time.Now().Add(au.expireDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(au.signingKey)
}
func (au *authUseCase) ParseToken(ctx context.Context, accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return au.signingKey, nil
	})
	if err != nil {
		return "", nil
	}
	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.UserId, nil
	}
	return "", auth.ErrInvalidAccessToken

}
