package dgraph

import (
	"context"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v2"
	dgo "github.com/dgraph-io/dgo/v2"
	api "github.com/dgraph-io/dgo/v2/protos/api"
	grpc "google.golang.org/grpc"
)

type Db struct {
	*dgo.Dgraph
}

func New() (*Db, error) {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(api.NewDgraphClient(d)), err
}

type Buyer struct {
	buyer_id string `json:"buyer_id,omitempty"`
	name     string `json:"name,omitempty"`
	age      int    `json:"age,omitempty"`
	date     string `json:"date,omitempty"`
}

type Product struct {
	product_id string `json:"product_id,omitempty"`
	name       string `json:"name,omitempty"`
	price      int    `json:"price,omitempty"`
	date       string `json:"date,omitempty"`
}

type Transaction struct {
	transaction_id string    `json:"transaction_id,omitempty"`
	buyer_id       string    `json:"buyer_id,omitempty"`
	ip             string    `json:"ip,omitempty"`
	device         string    `json:"device,omitempty"`
	products       []Product `json:"products,omitempty"`
	date           string    `json:"date,omitempty"`
}

func (d *Db) getBuyerById(buyer_id string) Buyer {
	txn := d.Dgraph.newReadOnlyTxn()
	resp, err := txn.Query(context.Background(), `{
		me(fun: eq(Buyer.buyer_id,buyer_id)){
			Buyer.buyer_id
			Buyer.name
			Buyer.age
			Buyer.date
		}
	}`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n", resp.Json)
	var b Buyer
	err = resp.Json.data.me
	b.buyer_id = err.Buyer.buyer_id
	b.name = err.Buyer.name
	b.age = err.Buyer.age
	b.date = err.Buyer.date

	return b
}
