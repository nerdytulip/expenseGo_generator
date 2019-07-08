package types

import(

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

func RouteHandler(){
  r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/expenses", func(r chi.Router) {
		r.Post("/", Create)
		r.Get("/", GetAll)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(Ctx)
			r.Get("/", GetOne)
			r.Put("/", Update)
			r.Delete("/", Delete)
		})
	})

	log.Fatal(http.ListenAndServe(":8080", r))

}