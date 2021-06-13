package dgraphql

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

type Db struct {
	*dgo.Dgraph
}

func New() (*Db, *grpc.ClientConn, error) {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	db := dgo.NewDgraphClient(api.NewDgraphClient(d))

	return &Db{db}, d, err
}

type Buyer struct {
	BuyerID string
	Name    string
	Age     int
	Date    string
}

type Product struct {
	ProductID string
	Name      string
	Price     int
	Date      string
}

type Transaction struct {
	TransactionID string
	BuyerID       string
	Ip            string
	Device        string
	Products      []Product
	Date          string
}

func (d *Db) GetBuyerById(buyer_id string) Buyer {
	txn := d.NewReadOnlyTxn()
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
	err = json.Unmarshal(resp.GetJson(), &b)

	if err != nil {
		log.Fatal(err)
	}

	return b
}
