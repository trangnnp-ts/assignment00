package router

import (
	"net/http"

	"github.com/go-chi/chi"
	user "github.com/trangnnp-ts/assignment00/controller"
)

func Routerdef() {
	routerx := chi.NewRouter()
	routerx.Post("/data", user.InputData)
	routerx.Get("/data/{alias}", user.GetOne)
	routerx.Get("/data", user.GetAll)
	routerx.Put("/data/{alias}", user.Update)
	routerx.Delete("/data/{alias}", user.Remove)
	http.ListenAndServe(":3000", routerx)
}
