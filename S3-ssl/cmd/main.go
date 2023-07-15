package main

import (
	"blocopad_secure/internal/db"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"blocopad_secure/internal/backend"

	"github.com/gorilla/mux"
)

// para rodar o redis
// docker run -p 6379:6379 --name some-redis -d redis

func WriteResponse(status int, body interface{}, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(body)
	w.Write(payload)
}

func ReadNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if data, err := backend.GetKey(id); err == nil {
		WriteResponse(200, data, w)
	} else {
		if err.Error() == "not found" {
			WriteResponse(404, "Note not found", w)
		} else {
			WriteResponse(500, "Error", w)
		}
	}
}

func WriteNote(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var note db.Note
	if err := decoder.Decode(&note); err != nil {
		WriteResponse(http.StatusBadRequest, map[string]string{"error": err.Error()}, w)
		return
	}
	uuidString, err := backend.SaveKey(note.Text, note.OneTime)
	if err != nil {
		WriteResponse(http.StatusBadRequest, map[string]string{"error": "invalid request"}, w)
	} else {
		WriteResponse(200, map[string]string{"code": uuidString}, w)
	}
}

func main() {
	serverPort := "8443"
	if port, hasValue := os.LookupEnv("API_PORT"); hasValue {
		serverPort = port
	}

	databaseUrl := "localhost:6379"
	if dbUrl, hasValue := os.LookupEnv("API_DB_URL"); hasValue {
		databaseUrl = dbUrl
	}

	databasePassword := ""
	if dbPassword, hasValue := os.LookupEnv("API_DB_PASSWORD"); hasValue {
		databasePassword = dbPassword
	}

	certificatePath := "./server.crt"
	if cPath, hasValue := os.LookupEnv("API_CERT_PATH"); hasValue {
		certificatePath = cPath
	}

	privateKeyPath := "./server.key"
	if pkPath, hasValue := os.LookupEnv("API_PK_PATH"); hasValue {
		privateKeyPath = pkPath
	}

	db.DatabaseUrl = databaseUrl
	db.DatabasePassword = databasePassword

	router := mux.NewRouter()
	router.HandleFunc("/api/note/{id}", ReadNote).Methods("GET")
	router.HandleFunc("/api/note", WriteNote).Methods("POST")
	err := http.ListenAndServeTLS(fmt.Sprintf(":%s", serverPort), certificatePath, privateKeyPath, router)
	fmt.Println(err)
}
