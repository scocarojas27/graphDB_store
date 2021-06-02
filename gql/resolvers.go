package gql

import (
	"github.com/graphql-go/graphql"
	dg "github.com/scocarojas27/graphDB_store/dgraphql"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	db *dg.Db
}

// UserResolver resolves our user query through a db call to GetUserByName
func (r *Resolver) BuyerResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	buyer_id, ok := p.Args["buyer_id"].(string)
	if ok {
		buyer := r.db.getBuyerById(buyer_id)
		return buyer, nil
	}

	return nil, nil
}
