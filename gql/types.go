package gql

import "github.com/graphql-go/graphql"

// User describes a graphql object containing a User
var Buyer = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Buyer",
		Fields: graphql.Fields{
			"buyer_id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"age": &graphql.Field{
				Type: graphql.Int,
			},
			"date": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
