package types

import(

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

var mh *MongoHandler

func RouteHandler(m *MongoHandler){
  r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/expenses", func(r chi.Router) {
		r.Post("/", m.Create)
		r.Get("/", m.GetAll)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(m.Ctx)
			r.Get("/", m.GetOne)
			r.Put("/", m.Update)
			r.Delete("/",m.Delete)
		})
	})

	log.Fatal(http.ListenAndServe(":8080", r))

}

func main() {
mongoDbConnection := "mongodb://localhost:27017"
//TODO implementation
mh = NewHandler(mongoDbConnection)
RouteHandler(mh )

}