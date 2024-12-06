package main

import (
	"log"
	"net/http"

	"Go_CRUD_server/router"
	"Go_CRUD_server/services"
	"Go_CRUD_server/utils"
)

func main() {
	log.Println("In my Main App")

	var dbconn = utils.GetConnection()
	services.SetDB(dbconn)
	var appRouter = router.CreateRouter()

	log.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
}
