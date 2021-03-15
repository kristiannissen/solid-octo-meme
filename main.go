package main

import (
	"log"
	"net/http"
	"os"
    "fmt"
    "encoding/json"
)

type Result struct {
    Message string
}

func readMessage(w http.ResponseWriter, r *http.Request) {
    var res Result

    err := json.NewDecoder(r.Body).Decode(&res)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }
    fmt.Println(res)
}

func main() {
	var port string = os.Getenv("PORT")

	if port == "" {
		port = "80"
	}

    mux := http.NewServeMux()
	mux.HandleFunc("/", readMessage)

    log.Println(fmt.Sprintf("Listening on port %s", port))

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
