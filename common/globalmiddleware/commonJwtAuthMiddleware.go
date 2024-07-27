package globalmiddleware

import (
	"MyDouSheng/common/ctxdata"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type CommonJwtAuthMiddleware struct {
	secret string
	Id     float64 `json:"user_id"`
	jwt.StandardClaims
}
type ctxKey string

const (
	ClaimsKey ctxKey = "claims"
)

func NewCommonJwtAuthMiddleware(secret string) *CommonJwtAuthMiddleware {
	return &CommonJwtAuthMiddleware{
		secret: secret,
	}
}

func (m *CommonJwtAuthMiddleware) Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(strings.NewReader(string(body)))
		var requestBody map[string]interface{}
		json.Unmarshal(body, &requestBody)
		if !strings.Contains(string(body), "token") {
			fmt.Println("没有token")
		} else {
			fmt.Println("含有token")
			tokenString, _ := requestBody["token"].(string)
			fmt.Println(tokenString)
			token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return []byte(m.secret), nil
			})
			if token == nil || !token.Valid {
				fmt.Println(-1)
				ctxWithClaims := context.WithValue(r.Context(), ClaimsKey, "-1")
				r = r.WithContext(ctxWithClaims)
			}
			if token != nil {
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					expTime := time.Unix(int64(claims["exp"].(float64)), 0)
					if expTime.Before(time.Now()) {
						fmt.Println(-1)
						ctxWithClaims := context.WithValue(r.Context(), ClaimsKey, "-1")
						r = r.WithContext(ctxWithClaims)
					} else {
						claims_userId, _ := claims[ctxdata.CtxKeyJwtUserId].(float64)
						ctxWithClaims := context.WithValue(r.Context(), ctxdata.CtxKeyJwtUserId, int64(claims_userId))
						fmt.Println("id", claims_userId)
						r = r.WithContext(ctxWithClaims)
					}
				}
			}
		}
		next(w, r)
	}
}
