package router

import (
	"net/http"
	"serverDemo/controller"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/user/login", controller.Login)

	return mux
}
