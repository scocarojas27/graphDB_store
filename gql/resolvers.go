package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/scocarojas27/graphDB_store/dgraphql"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	db *dgraphql.Db
}

// BuyerResolver resolves our user query through a db call to GetBuyerById
func (r *Resolver) BuyerResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	buyer_id, ok := p.Args["BuyerID"].(string)
	if ok {
		buyer := r.db.GetBuyerById(buyer_id)
		return buyer, nil
	}
	return nil, nil
}

// BuyersResolver resolves our user query through a db call to GetAllBuyers
func (r *Resolver) BuyersResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	buyer := r.db.GetAllBuyers()
	return buyer, nil
}

// InsertProductsResolver resolves our user query through a db call to InsertProducts
func (r *Resolver) InsertProductsResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	products := r.db.InsertProducts()
	return products, nil
}

// InsertBuyersResolver resolves our user query through a db call to InsertBuyers
func (r *Resolver) InsertBuyersResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	buyers := r.db.InsertBuyers()
	return buyers, nil
}

// ReportResolver resolves our user query through a db call to GetReport
func (r *Resolver) ReportResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	buyer_id, ok := p.Args["BuyerID"].(string)
	if ok {
		report := r.db.GetReport(buyer_id)
		return report, nil
	}
	return nil, nil
}

// ProductResolver resolves our user query through a db call to GetProductById
func (r *Resolver) ProductResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	product_id, ok := p.Args["ProductID"].(string)
	if ok {
		product, _ := r.db.GetProductById(product_id)
		return product, nil
	}
	return nil, nil
}

// BuyerTransactionsResolver resolves our user query through a db call to GetTransactionsByBuyerById
func (r *Resolver) BuyerTransactionResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	buyer_id, ok := p.Args["BuyerID"].(string)
	if ok {
		transactions := r.db.GetTransactionsByBuyerId(buyer_id)
		return transactions, nil
	}
	return nil, nil
}
