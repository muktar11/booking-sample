package main 



import (
	"fmt"
	"net/http"
)


func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}
//NoSurd adds CSRF protection to a;; POST request
func NoSurf(next http.Handler)http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCokkie(http.Cookie{
		HttpOnly: true,
		Path:"/", 
		Secure: false, 
		SameSite: http.SameSiteLaxMode
	})
	return csrfHandler 
}
//SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}