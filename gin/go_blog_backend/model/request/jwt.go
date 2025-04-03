package request

import (
	"blog_backend/model/appTypes"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UserID uint
	UUID   uuid.UUID
	RoleID appTypes.RoleID
}
