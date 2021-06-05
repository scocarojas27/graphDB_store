package gql

import (
	"github.com/dgraph-io/dgo/v2"
	"github.com/graphql-go/graphql"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	db *dgo.Dgraph
}

// UserResolver resolves our user query through a db call to GetUserByName
func (r *Resolver) BuyerResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	buyer_id, ok := p.Args["buyer_id"].(string)
	if ok {
		buyer := dg.getBuyerById(buyer_id, r.db)
		return buyer, nil
	}

	return nil, nil
}
