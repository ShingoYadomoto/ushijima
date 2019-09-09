package middleware

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

const (
	basicAuthUser     = "user"
	basicAuthPassword = "password"
)

func BasicAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			w := c.Response().Writer
			r := c.Request()

			// Basic認証のヘッダ解析に失敗したらokにfalseが入るみたい
			user, pass, ok := r.BasicAuth()
			// 入力された内容のチェック
			if ok == false || user != basicAuthUser || pass != basicAuthPassword {
				w.Header().Set("WWW-Authenticate", `Basic realm="auth area"`)
				http.Error(w, "needs authenticate", http.StatusUnauthorized)
				return nil
			}

			log.Debug("Basic auth passed.")
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}
