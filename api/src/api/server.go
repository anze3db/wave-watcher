package api

import (
	"github.com/ant0ine/go-json-rest/rest"
	//	"gofetch"
	"log"
	"net/http"
)

// Serve something
func Serve() {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/updates", getAllUpdates),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Printf("heello")
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", api.MakeHandler()))
}

//
// func getAllUpdates(w rest.ResponseWriter, r *rest.Request) {
// 	updates := []gofetch.Update{}
// 	session := gofetch.DB.Session()
// 	session.Find(&updates)
// 	w.WriteJson(&updates)
// }

func getAllUpdates(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(1)
}
