package main

import (
	"flag"
	"fmt"
	"github.com/Jackalgit/Gofermat/cmd/config"
	"github.com/Jackalgit/Gofermat/internal/database"
	"github.com/Jackalgit/Gofermat/internal/handlers"
	"github.com/Jackalgit/Gofermat/internal/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func init() {
	config.ConfigServerPort()
	config.ConfigLogger()
	config.ConfigDatabaseDSN()
	config.ConfigAccrualSystem()
	config.ConfigSecretKey()
}

func main() {

	flag.Parse()
	fmt.Println(time.Now().Format(time.RFC3339))

	if err := runServer(); err != nil {
		log.Println("runServer ERROR: ", err)
	}

}

func runServer() error {

	handler := &handlers.GoferMat{
		Storage:         database.NewDataBase(),
		DictUserIDToken: models.NewDictUserIDToken(),
	}

	router := mux.NewRouter()

	router.HandleFunc("/ping", handler.PingDB).Methods("GET")
	router.HandleFunc("/api/user/register", handler.Register).Methods("POST")
	router.HandleFunc("/api/user/login", handler.Login).Methods("POST")
	router.HandleFunc("/api/user/orders", handler.ListOrders).Methods("GET", "POST")
	router.HandleFunc("/api/user/balance", handler.Balance).Methods("GET")
	router.HandleFunc("/api/user/balance/withdraw", handler.Withdraw).Methods("POST")
	router.HandleFunc("/api/user/withdrawals", handler.Withdrawals).Methods("GET")

	if err := http.ListenAndServe(config.Config.ServerPort, router); err != nil {
		return fmt.Errorf("[ListenAndServe] запустить сервер: %q", err)

	}

	return http.ListenAndServe(config.Config.ServerPort, router)

}