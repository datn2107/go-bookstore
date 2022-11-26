package main

import (
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/datn2107/go-bookstore/pkg/routers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routers.RegisterBookStoreRouters(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
