package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/scocarojas27/graphDB_store/dgraphql"
)

// Root holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot(db *dgraphql.Db) *Root {
	// Create a resolver holding our databse. Resolver can be found in resolvers.go
	resolver := Resolver{db: db}

	// Create a new Root that describes our base query set up. In this
	// example we have a user query that takes one argument called name
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"buyer": &graphql.Field{
						// Slice of User type which can be found in types.go
						Type: graphql.NewList(Buyer),
						Args: graphql.FieldConfigArgument{
							"buyer_id": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.BuyerResolver,
					},
				},
			},
		),
	}
	return &root
}
