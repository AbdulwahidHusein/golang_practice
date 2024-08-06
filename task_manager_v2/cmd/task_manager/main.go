package main

//"github.com/gin-gonic/gin"
import (
	"net/http"
	"task_managemet_api/cmd/task_manager/pkg/db"

	"context"
	"log"

	"task_managemet_api/cmd/task_manager/cmd/bootstrap"
)

func main() {
	client, err := db.ConnectToMongo()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	router := bootstrap.GetRouter(client)

	http.ListenAndServe(":8080", router)
}
