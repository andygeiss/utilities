package safe

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// HandleRequest ...
func HandleRequest(r *http.Request) (start time.Time, out interface{}) {
	start = time.Now()
	if err := json.NewDecoder(r.Body).Decode(&out); err != nil {
		log.Println(err)
	}
	log.Printf("[%-8s] [%-20s] -> [%v]", r.Method, r.RequestURI, out)
	return start, out
}

// HandleResponse ...
func HandleResponse(w http.ResponseWriter, r *http.Request, data interface{}, start time.Time) {
	if err := json.NewEncoder(w).Encode(&data); err != nil {
		log.Println(err)
	}
	log.Printf("[%-8s] [%-20s] <- [%v]", r.Method, r.RequestURI, data)
	log.Printf("[%-8s] [%-20s] response time [%v]", r.Method, r.RequestURI, time.Since(start))
}
