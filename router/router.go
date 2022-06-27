package router

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"task-assign/handlers"
)

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff here
	fmt.Println("-----------------------------------------")
        fmt.Println("Device: ",r.UserAgent(),"\nFrom:   ",r.RemoteAddr,"\nMethods:",r.Method,"\nRoute:  ",r.RequestURI)
        // Call the next handler, which can be another middleware in the chain, or the final handler.
	fmt.Println("-----------------------------------------")
        next.ServeHTTP(w, r)
    })
}

//allowJson() is used by routes that are supposed
//to send json data. It checks if the content-type
// is set to application/json
func allowJson(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if r.Header.Get("Content-type") != "application/json"{
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type. Only JSON files are allowed\n"))
			return
		}

		next.ServeHTTP(w,r)
	})
}

var Rt = mux.NewRouter()

func SetUpRoutes() {
	//jsonPath routes
	jsonPath := Rt.PathPrefix("/data").Subrouter()
	
	//auth routes
	auth := jsonPath.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register/",handlers.Registration).Methods("POST")
	auth.HandleFunc("/login/",handlers.Login).Methods("POST")
	
	
	//stuff routes
	r := Rt.PathPrefix("/rr").Subrouter()
	r.HandleFunc("/", handlers.TestFunc).Methods("GET")
	
	//assign middlewares
	Rt.Use(loggingMiddleware)
	jsonPath.Use(allowJson)
}
