package routes

import (
	"net/http"

	"../controllers"
	"github.com/gorilla/mux"
	// "fmt"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("GET", "/findImage_all", controllers.FindImage_all, nil)
	register("GET", "/findImage_by_tag", controllers.FindImage_by_tag, nil)
	register("GET", "/newpost", controllers.NewPost, nil)
	register("POST", "/createImage", controllers.CreateImage, nil)
	register("GET", "/deleteImage", controllers.DeleteImage, nil)
	register("GET", "/findImageNum", controllers.FindImageNum, nil)
	register("GET", "/addFavorite", controllers.AddFavorite, nil)
	register("GET", "/addLocation", controllers.AddLocation, nil)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		r := router.Methods(route.Method).
			Path(route.Pattern)
		if route.Pattern == "" {
			r = router.Methods(route.Method).
				PathPrefix(route.Pattern)
		}
		if route.Middleware != nil {
			r.Handler(route.Middleware(route.Handler))
		} else {
			r.Handler(route.Handler)
		}
	}
	return router
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
