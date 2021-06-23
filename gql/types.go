package gql

import "github.com/graphql-go/graphql"

// Buyer describes a graphql object containing a buyer
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

// Product describes a graphql object containing a product
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

// Report describes a graphql object containing a buyer report
var Report = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Report",
		Fields: graphql.Fields{
			"Transactions": &graphql.Field{
				Type: &graphql.List{
					OfType: graphql.String,
				},
			},
			"SameIp": &graphql.Field{
				Type: &graphql.List{
					OfType: graphql.String,
				},
			},
			"Recomendations": &graphql.Field{
				Type: &graphql.List{
					OfType: graphql.String,
				},
			},
		},
	},
)

// Transaction describes a graphql object containing a transaction
var Transaction = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Transaction",
		Fields: graphql.Fields{
			"TransactionID": &graphql.Field{
				Type: graphql.String,
			},
			"BuyerID": &graphql.Field{
				Type: graphql.String,
			},
			"Ip": &graphql.Field{
				Type: graphql.String,
			},
			"Device": &graphql.Field{
				Type: graphql.String,
			},
			"Products": &graphql.Field{
				Type: &graphql.List{
					OfType: graphql.String,
				},
			},
			"Date": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
