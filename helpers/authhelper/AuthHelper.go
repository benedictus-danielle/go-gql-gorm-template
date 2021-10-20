package authhelper

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"gorm.io/gorm"
)

// type contextKey struct {
// 	name string
// }
//
// var userCtxKey = &contextKey{"user"}

func Middleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}
			tokens, claims, err := jwtauth.FromContext(r.Context())
			if err != nil || tokens == nil {
				w.WriteHeader(401)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"success": false,
					"data":    "Unauthorized",
				})
				return
			}
			if time.Now().After(claims["exp"].(time.Time)) {
				w.WriteHeader(401)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"success": false,
					"data":    "Token Expired",
				})
				return
			}
			//! Uncomment this section when there are user implemented
			//var user model.User
			//err = db.Preload("Role.MappingPermissions.Permission").Where("id = ?", claims["sub"].(string)).First(&user).Error
			//if err != nil {
			//	w.WriteHeader(401)
			//	json.NewEncoder(w).Encode(map[string]interface{}{
			//		"success": false,
			//		"data":    "User doesn't exists",
			//	})
			//}
			//ctx := context.WithValue(r.Context(), userCtxKey, &user)
			//r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

//! Uncomment this section when there are user implemented
//func ForContext(ctx context.Context) *model.User {
//	raw, _ := ctx.Value(userCtxKey).(*model.User)
//	return raw
//}
//
//func Authenticate(user *model.User) *gqlerror.Error {
//	if user == nil {
//		return gqlhelper.CreateError(
//			jsonhelper.SerializeMap(
//				map[string]interface{}{
//					"success": false,
//					"status":  401,
//					"message": "Unauthorized",
//				},
//			),
//		)
//	}
//	return nil
//}
