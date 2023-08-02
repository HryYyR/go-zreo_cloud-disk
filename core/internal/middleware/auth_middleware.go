package middleware

import (
	"cloud_disk/go-zreo_cloud-disk/core/helper"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			logx.Errorv("Invalid authorization")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid authorization"))
			return
		}
		uc, err := helper.DecryptToekn(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid authorization: " + err.Error()))
			return
		}

		r.Header.Set("UserId", uc.ID)
		r.Header.Set("UserIdentity", uc.Identity)
		r.Header.Set("UserName", uc.Name)
		next(w, r)
	}
}
