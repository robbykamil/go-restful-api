// 3. middleware funtion, it used to wrap HTTP handlers
// and add additional functionality such as logging, authentication, rate limit, etc.

package main

import(
	"log"
	"net/http" //provide http server and client functionalities
	"time"
)

func Logger(inner http.Handler, name string) http.Handler { // http.Handler is main logic to handling requests
	// convert provided function into http.Handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//measure how long the request took
		start := time.Now()

		//processing actual the http request
		inner.ServeHTTP(w, r)

		//logging the request
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,			//HTTP method e.g. GET
			r.RequestURI,		//The request URI e.g. "/todos"
			name,				// Name of the handler
			time.Since(start),  // time taken to process the request
		)
	})
}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         