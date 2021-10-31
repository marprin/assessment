package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/marprin/assessment/fetchapp/pkg/jwt"
	"github.com/marprin/assessment/fetchapp/pkg/response"
	"github.com/sirupsen/logrus"
)

const (
	NameCtx  = "name"
	PhoneCtx = "phone"
	RoleCtx  = "role"
)

type (
	Middleware struct {
		jwtRepo jwt.Repository
	}
	Options struct {
		JwtRepo jwt.Repository
	}
)

func New(o *Options) *Middleware {
	return &Middleware{
		jwtRepo: o.JwtRepo,
	}
}

func (m *Middleware) Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"Error":  err,
					"Method": r.Method,
					"Path":   r.URL.Path,
				}).Errorln("Panic Interceptor")

				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}

		}()

		next.ServeHTTP(w, r)

	})
}

func (m *Middleware) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		authorizationToken := r.Header.Get("Authorization")
		if authorizationToken == "" {
			response.Response(w, r, http.StatusUnauthorized, nil, nil, nil, http.StatusUnauthorized)
			return
		}

		token := strings.Replace(authorizationToken, "Bearer ", "", 1)
		tokenData, err := m.jwtRepo.ExtractToken(r.Context(), token)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err":   err,
				"token": token,
			}).Errorln("[Authorize] Error when validate token")
			response.Response(w, r, http.StatusUnauthorized, nil, nil, nil, http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, PhoneCtx, tokenData.Phone)
		ctx = context.WithValue(ctx, RoleCtx, tokenData.Role)
		ctx = context.WithValue(ctx, NameCtx, tokenData.Name)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
