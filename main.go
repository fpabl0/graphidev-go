package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func parseHeaderString(s string) map[string]string {
	m := map[string]string{}
	if s == "" {
		return nil
	}
	headers := strings.Split(s, ",")
	for _, h := range headers {
		parts := strings.Split(h, ":")
		if len(parts) != 2 {
			panic("invalid header fmt, must be of the form <header_name1>:<header_value1>,<header_name2>:<header_value2>")
		}
		m[parts[0]] = parts[1]
	}
	return m
}

func main() {
	var uiPort, gqlPort int
	var headerString string
	flag.IntVar(&uiPort, "uiport", -1, "The port where the web ui will host the graphiql interface")
	flag.IntVar(&gqlPort, "gqlport", -1, "The port where the graphql server is served")
	flag.StringVar(&headerString, "H", "", "Custom HTTP Headers written in this form <header_name1>:<header_value1>,<header_name2>:<header_value2>")
	flag.Parse()

	if uiPort < 0 {
		uiPort = 3001
		log.Printf("[WARN]  uiport not set (using default: %d)\n", uiPort)
	}

	if gqlPort < 0 {
		gqlPort = 3000
		log.Printf("[WARN]  gqlport not set (using default: %d)\n", gqlPort)
	}

	if headerString == "" {
		log.Printf("[WARN]  http headers empty")
	}

	h := GraphiQL{
		GraphqlURL: fmt.Sprintf("http://localhost:%d/graphql", gqlPort),
		HeaderMap:  parseHeaderString(headerString),
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
