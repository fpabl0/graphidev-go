package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var uiPort, gqlPort int
	flag.IntVar(&uiPort, "uiport", -1, "The port where the web ui will host the graphiql interface")
	flag.IntVar(&gqlPort, "gqlport", -1, "The port where the graphql server is served")
	flag.Parse()

	if uiPort < 0 {
		uiPort = 3001
		log.Printf("[WARN]  uiport not set (using default: %d)\n", uiPort)
	}

	if gqlPort < 0 {
		gqlPort = 3000
		log.Printf("[WARN]  gqlport not set (using default: %d)\n", gqlPort)
	}

	h := GraphiQL{
		GraphqlURL: fmt.Sprintf("http://localhost:%d/graphql", gqlPort),
	}

	mux := http.NewServeMux()
	mux.Handle("/", h)

	s := &http.Server{Addr: fmt.Sprintf("localhost:%d", uiPort), Handler: mux}

	fmt.Println()
	log.Printf("[INFO]  Browse to %s in order to use graphiql interface\n", s.Addr)

	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}

	log.Println("shut down")
}
