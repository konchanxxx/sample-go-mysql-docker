package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	dbDialect = "mysql"
)

func init() {
	conn, err := getConn()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.AutoMigrate(&Product{})
}

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

// Product は商品
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	conn, err := getConn()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	err = conn.Create(&Product{Name: "switch", Price: 29980}).Error
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func getProductLatest(w http.ResponseWriter, r *http.Request) {
	conn, err := getConn()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	p := &Product{}
	err = conn.Last(p).Error
	if err != nil {
		panic(err)
	}

	res, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func getConn() (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4",
		"user",
		"password",
		"db",
		"3306",
		"sample",
	)

	return gorm.Open(dbDialect, connectionString)
}
