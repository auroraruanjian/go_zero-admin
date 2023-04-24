package middleware

import "net/http"

type CheckUrlMiddleware struct {
}

func NewCheckUrlMiddleware() *CheckUrlMiddleware {
	return &CheckUrlMiddleware{}
}

func (m *CheckUrlMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
