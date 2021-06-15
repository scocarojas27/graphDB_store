package gql

import "github.com/graphql-go/graphql"

// User describes a graphql object containing a User
var Buyer = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Buyer",
		Fields: graphql.Fields{
			"BuyerID": &graphql.Field{
				Type: graphql.String,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Age": &graphql.Field{
				Type: graphql.Int,
			},
			"Date": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var Product = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"ProductID": &graphql.Field{
				Type: graphql.String,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Price": &graphql.Field{
				Type: graphql.Int,
			},
			"Date": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
