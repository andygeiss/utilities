package logging

import (
	"log"
	"net/http"
	"time"
)

// Response ...
func Response(r *http.Request, rid string, response interface{}, err error, start time.Time) {
	log.Printf("[%-8s] [%-20s] [%s] response [%p] round trip [%v] error [%v]", r.Method, r.RequestURI, rid, response, time.Since(start), err)
}
