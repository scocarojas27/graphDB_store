package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/scocarojas27/graphDB_store/dgraph"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	db *dgraph.Db
}

// UserResolver resolves our user query through a db call to GetUserByName
func (r *Resolver) BuyerResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	buyer_id, ok := p.Args["buyer_id"].(string)
	if ok {
		users := r.db.getBuyerById(buyer_id)
		return users, nil
	}

	return nil, nil
}
