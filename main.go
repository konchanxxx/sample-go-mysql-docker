package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// 商品を追加するendpoint
	router.HandleFunc("/products", createProduct).Methods("POST")
	// 最新の商品情報を取得するendpoint
	router.HandleFunc("/products/latest", getProductLatest).Methods("GET")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	// TODO 商品を追加するような処理をかく
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("success create product"))
}

func getProductLatest(w http.ResponseWriter, r *http.Request) {
	// TODO 最新の商品を取得するような処理をかく
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("success get product latest"))
}
