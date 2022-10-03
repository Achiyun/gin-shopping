package request

import (
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	jwt.StandardClaims
}

type BaseClaims struct {
	ID       int
	UUID     uuid.UUID
	Username string
	// 角色id
	RoleId uint
}
