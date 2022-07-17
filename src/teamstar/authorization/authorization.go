package authorization

import (
	"errors"
	"net/http"

	"github.com/alexedwards/scs/v2"
	log "github.com/sirupsen/logrus"

	"github.com/casbin/casbin"

	"github.com/MarcBeckRT/myapp-go/src/teamstar/model"
	"github.com/MarcBeckRT/myapp-go/src/teamstar/service"
)

func Authorizer(e *casbin.Enforcer, users *model.User) func(next http.Handler) http.Handler {
	var sessionManager *scs.SessionManager
	return func(next http.Handler) http.Handler {
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
					log.Errorf("403 FORBIDDEN", w, errors.New("user does not exist"))
					return
				}
			}
			// casbin enforce
			res, err := e.EnforceSafe(role, r.URL.Path, r.Method)
			if err != nil {
				log.Errorf("500 ERROR", w, err)
				return
			}
			if res {
				next.ServeHTTP(w, r)
			} else {
				log.Errorf("403 FORBIDDEN", w, errors.New("unauthorized"))
				return
			}
		}

		return http.HandlerFunc(fn)
	}
}
