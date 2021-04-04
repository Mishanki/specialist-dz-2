package main

import (
	"fmt"
	"github.com/Mishanki/specialist-dz-2/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var (
	prefix = "/api/v1/"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc(prefix+"register", handlers.Register).Methods("POST")

	// Возвращает JWT метку для зарегестрированных пользователей.
	router.HandleFunc(prefix+"auth", handlers.Auth).Methods("GET")

	// Методы требующие JWT токены

	// GET /auto/<string:mark> - возвращает информацию про автомобиль с именем mark и код 200. В случае, если автомобиля нет в БД в текущий момент возвращаем {"Error" : "Auto with that mark not found"} и код 404.
	router.HandleFunc(prefix+"auto/{mark}", handlers.GetAuto).Methods("GET")

	// POST /auto/<string:mark> - добавляет автомобиль с именем mark в БД. В случае успеха - 201 и сообщение {"Message" : "Auto created"}. В случае, если автомобиль с таким именем уже существует - 400 и {"Error" : "Auto with that mark exists"}. Структура автомобиля в теле запроса выглядит следующим образом:
	router.HandleFunc(prefix+"auto/{mark}", handlers.CreateAuto).Methods("POST")

	// PUT /auto/<string:mark> - обновляет информацию про автомобиль с именем mark в БД. В случае успеха - 202 и сообщение {"Message" : "Auto updated"}. В случае, если автомобиля нет в БД в текущий момент возвращаем {"Error" : "Auto with that mark not found"} и код 404.
	router.HandleFunc(prefix+"auto/{mark}", handlers.UpdateAuto).Methods("PUT")

	// DELETE /auto/<string:mark> - удаляет информацию про автомобиль с именем mark из БД. В случае успеха - 202 и сообщение {"Message" : "Auto deleted"}. В случае, если автомобиля нет в БД в текущий момент возвращаем {"Error" : "Auto with that mark not found"} и код 404.
	router.HandleFunc(prefix+"/auto/mark", handlers.DeleteAuto).Methods("DELETE")

	// GET /stock - возвращает информацию про все имеющиеся на данный момент в БД автомобили и код 200 в случае, если имеется хотя бы один автомобиль в наличии. В противном случае - 400 и сообщение {"Error" : "No one autos found in DataBase"}.
	router.HandleFunc(prefix+"/stock", handlers.GetStock).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		panic("Param PORT is not found in .env")
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
