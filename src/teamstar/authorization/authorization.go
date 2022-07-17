package authorization

import (
	"errors"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/casbin/casbin"

	"github.com/MarcBeckRT/myapp-go/src/teamstar/model"
	"github.com/MarcBeckRT/myapp-go/src/teamstar/service"
)

// Authorizer is a middleware for authorization
func Authorizer(e *casbin.Enforcer, users *model.User) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		var sessionManager = scs.New()
		fn := func(w http.ResponseWriter, r *http.Request) {
			role := sessionManager.GetString(r.Context(), "role")
			if role == "" {
				role = "anonymous"
			}
			// if it's a player, check if the user still exists
			if role == "player" {
				uid := sessionManager.GetInt(r.Context(), "userID")
				exists := service.Exists(uid)
				if !exists {
					writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("user does not exist"))
					return
				}
			}
			// casbin enforce
			res, err := e.EnforceSafe(role, r.URL.Path, r.Method)
			if err != nil {
				writeError(http.StatusInternalServerError, "ERROR", w, err)
				return
			}
			if res {
				next.ServeHTTP(w, r)
			} else {
				writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("unauthorized"))
				return
			}
		}

		return http.HandlerFunc(fn)
	}
}

func writeError(status int, message string, w http.ResponseWriter, err error) {
	log.Print("ERROR: ", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(message))
}
