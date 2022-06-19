package app

import (
	"github.com/rmortale/golang-microservices/mvc/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
