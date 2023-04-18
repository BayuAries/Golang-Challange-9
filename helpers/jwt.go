package helpers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "rahasia"

func GenerateToken(id uint, email string, level string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"level": level,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResp := errors.New("filed to verify token")

	headerToken := c.Request.Header.Get("Authorization")
	isBearer := strings.HasPrefix(headerToken, "Bearer")

	if !isBearer {
		return nil, errResp
	}

	// bearer : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJheUBnbWFpbC5jb20iLCJpZCI6MX0.QF3vrFM_Sh53lsEPQx_LHa2lg3vIgSzvAzx3sXzCxsU
	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResp
		}

		return []byte(secretKey), nil
	})

	//validate
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errResp
	}

	return token.Claims.(jwt.MapClaims), nil
}
