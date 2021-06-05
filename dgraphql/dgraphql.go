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

/*type Db struct {
	*dgo.Dgraph
}*/

func New() (*dgo.Dgraph, error) {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db := dgo.NewDgraphClient(api.NewDgraphClient(d))

	return db, err
}

type Buyer struct {
	buyer_id string
	name     string
	age      int
	date     string
}

type Product struct {
	product_id string
	name       string
	price      int
	date       string
}

type Transaction struct {
	transaction_id string
	buyer_id       string
	ip             string
	device         string
	products       []Product
	date           string
}

func getBuyerById(buyer_id string, d *dgo.Dgraph) Buyer {
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
