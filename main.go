package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"github.com/scocarojas27/graphDB_store/dgraph"
	"github.com/scocarojas27/graphDB_store/gql"
	"github.com/scocarojas27/graphDB_store/server"
)

type Buyer struct {
	buyer_id string `json:"buyer_id,omitempty"`
	name     string `json:"name,omitempty"`
	age      int    `json:"age,omitempty"`
}

type Product struct {
	product_id string `json:"product_id,omitempty"`
	name       string `json:"name,omitempty"`
	price      int    `json:"price,omitempty"`
}

type Transaction struct {
	transaction_id string    `json:"transaction_id,omitempty"`
	buyer          Buyer     `json:"buyer,omitempty"`
	ip             string    `json:"ip,omitempty"`
	device         string    `json:"device,omitempty"`
	products       []Product `json:"products,omitempty"`
}

func main() {
	router, db := initializeAPI()
	defer db.Close()

	// Listen on port 4000 and if there's an error log it and exit
	log.Fatal(http.ListenAndServe(":4000", router))
}

func initializeAPI() (*chi.Mux, *dgraph.Db) {
	// Create a new router
	router := chi.NewRouter()

	// Create a new connection to our pg database
	db, err := dgraph.New()
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
		middleware.Logger,          // log api request calls
		middleware.DefaultCompress, // compress results, mostly gzipping assets and json
		middleware.StripSlashes,    // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,       // recover from panics without crashing server
	)

	// Create the graphql route with a Server method to handle it
	router.Post("/graphql", s.GraphQL())

	return router, db
}
