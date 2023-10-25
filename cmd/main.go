package main

import (
	"CRUD-GO/cmd/server/handler"
	"CRUD-GO/internal/product"
	"CRUD-GO/pkg/store"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err:= sql.Open("mysql", "root:Servidor2$@tcp(localhost:3306)/go_db")
	if err!= nil{
		panic(err)
	}

	errPing := db.Ping()
	if errPing != nil{
		panic(errPing)
	}

	storage := store.SqlStoreImp{DB: db}
	repo := product.RepositoryImpl{Store: &storage}
	service := product.Service{Repository: &repo}
	handler := handler.ProductHandler{ProdService: &service}

	r := gin.Default()
	r.GET("/products/:id", handler.GetById)

	r.Run()
}