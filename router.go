package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	ctx := context.Background()
	svc := new(frontendServer)
	r := mux.NewRouter()
	r.HandleFunc("/", svc.homeHandler).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	mustConnGRPC(ctx, &svc.productCatalogSvcConn, svc.productCatalogSvcAddr)
	return r
}
