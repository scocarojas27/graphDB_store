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

	// Create a new Root that describes our base query set up. It includes Buyer
	// to get a buyer, Buyers to get all the buyers, InsertProducts to insert all present
	//day's products, InsertBuyers to insert all present day's buyers, and so on...
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"Buyer": &graphql.Field{
						// Buyer type which can be found in types.go
						Type: graphql.NewNonNull(Buyer),
						Args: graphql.FieldConfigArgument{
							"BuyerID": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.BuyerResolver,
					},
					"Buyers": &graphql.Field{
						Type:    graphql.NewList(Buyer),
						Resolve: resolver.BuyersResolver,
					},
					"InsertProducts": &graphql.Field{
						Type:    graphql.NewNonNull(Product),
						Resolve: resolver.InsertProductsResolver,
					},
					"InsertBuyers": &graphql.Field{
						Type:    graphql.NewNonNull(Buyer),
						Resolve: resolver.InsertBuyersResolver,
					},
					"BuyerReport": &graphql.Field{
						Type: graphql.NewNonNull(Report),
						Args: graphql.FieldConfigArgument{
							"BuyerID": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.ReportResolver,
					},
					"Product": &graphql.Field{
						Type: graphql.NewNonNull(Product),
						Args: graphql.FieldConfigArgument{
							"ProductID": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.ProductResolver,
					},
					"BuyerTransactions": &graphql.Field{
						Type: graphql.NewList(Transaction),
						Args: graphql.FieldConfigArgument{
							"BuyerID": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.BuyerTransactionResolver,
					},
				},
			},
		),
	}
	return &root
}
