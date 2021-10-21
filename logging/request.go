package logging

import (
	"log"
	"net/http"
)

// Request ...
func Request(r *http.Request, rid string, request interface{}, err error) {
	log.Printf("[%-8s] [%-20s] [%s] request [%p] error [%v]", r.Method, r.RequestURI, rid, request, err)
}
