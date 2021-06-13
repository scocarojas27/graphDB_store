package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"github.com/scocarojas27/graphDB_store/dgraphql"
	"github.com/scocarojas27/graphDB_store/gql"
	"github.com/scocarojas27/graphDB_store/server"
	"google.golang.org/grpc"
)

type Buyer struct {
	buyer_id string
	name     string
	age      int
}

type Product struct {
	product_id string
	name       string
	price      int
}

type Transaction struct {
	transaction_id string
	buyer          Buyer
	ip             string
	device         string
	products       []Product
}

func main() {
	router, db, conn := initializeAPI()
	defer conn.Close()

	log.Fatal("DB info:", db)

	// Listen on port 4000 and if there's an error log it and exit
	fmt.Println(router)
	log.Fatal(http.ListenAndServe(":4000", router))
}

func initializeAPI() (*chi.Mux, *dgraphql.Db, *grpc.ClientConn) {
	// Create a new router
	router := chi.NewRouter()

	// Create a new connection to our pg database
	db, conn, err := dgraphql.New()
	if err != nil {
		log.Fatal(err)
	}

	// Create our root query for graphql
	rootQuery := gql.NewRoot(db)
	// Create a new graphql schema, passing in the the root query
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	// Create a server struct that holds a pointer to our database as well
	// as the address of our graphql schema
	s := server.Server{
		GqlSchema: &sc,
	}

	// Add some middleware to our router

	router.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,              // log api request calls
		middleware.Compress(5, "gzip"), // compress results, mostly gzipping assets and json
		middleware.StripSlashes,        // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,           // recover from panics without crashing server
	)

	// Create the graphql route with a Server method to handle it
	router.Post("/graphql", s.GraphQL())

	return router, db, conn
}
