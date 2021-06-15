package dgraphql

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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
	BuyerID string `json:"BuyerID,omitempty"`
	Name    string `json:"Name,omitempty"`
	Age     int    `json:"Age,omitempty"`
	Date    string `json:"Date,omitempty"`
}

type Product struct {
	ProductID string `json:"ProductID,omitempty"`
	Name      string `json:"Name,omitempty"`
	Price     int    `json:"Price,omitempty"`
	Date      string `json:"Date,omitempty"`
}

type Transaction struct {
	TransactionID string    `json:"TransactionID,omitempty"`
	BuyerID       string    `json:"BuyerID,omitempty"`
	Ip            string    `json:"Ip,omitempty"`
	Device        string    `json:"Device,omitempty"`
	Products      []Product `json:"Products,omitempty"`
	Date          string    `json:"Date,omitempty"`
}

func (d *Db) InsertProducts() []Product {

	response, err := http.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products")
	txn := d.NewTxn()

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	lines := bytes.NewReader(responseData)
	buffReader := bufio.NewReader(lines)
	scanner := bufio.NewScanner(buffReader)
	insertDate := string(strconv.FormatInt(time.Now().Unix(), 10))
	con := 0

	type RootInsertProducts struct {
		Me []Product `json:"me"`
	}

	var p []Product

	for scanner.Scan() {
		var r RootInsertProducts
		var pr Product
		line := scanner.Text()
		prueba := strings.Split(string(line), "'")
		uid := "_:Product" + strconv.Itoa(con) + insertDate
		price, _ := strconv.Atoi(prueba[2])
		variables := map[string]interface{}{"uid": uid, "ProductID": prueba[0], "Name": prueba[1], "Price": price, "Date": insertDate, "dgraph.type": "Product"}
		productJson, err := json.Marshal(variables)

		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(productJson, &r)

		if err != nil {
			log.Fatal(err)
		}

		pr.ProductID = "ssss"
		pr.Name = "sssss"
		pr.Price = 1400
		pr.Date = "sssss"

		p = append(p, pr)

		resp, err := txn.Mutate(context.Background(), &api.Mutation{SetJson: productJson})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Response: %s\n", resp.Json)
		con = con + 1
	}

	err = txn.Commit(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	return p

}

func (d *Db) InsertBuyers() []Buyer {

	response, err := http.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers")
	txn := d.NewTxn()

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	lines := bytes.NewReader(responseData)
	buffReader := bufio.NewReader(lines)
	scanner := bufio.NewScanner(buffReader)
	insertDate := string(strconv.FormatInt(time.Now().Unix(), 10))
	con := 0

	type RootInsertBuyers struct {
		Me []Buyer `json:"me"`
	}

	var b []Buyer

	for scanner.Scan() {
		var r RootInsertBuyers
		var bu Buyer
		line := scanner.Text()
		prueba := strings.Split(string(line), "'")
		uid := "_:Buyer" + strconv.Itoa(con) + insertDate
		age, _ := strconv.Atoi(prueba[2])
		variables := map[string]interface{}{"uid": uid, "BuyerID": prueba[0], "Name": prueba[1], "Age": age, "Date": insertDate, "dgraph.type": "Buyer"}
		productJson, err := json.Marshal(variables)

		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(productJson, &r)

		if err != nil {
			log.Fatal(err)
		}

		bu.BuyerID = "ssss"
		bu.Name = "sssss"
		bu.Age = 14
		bu.Date = "sssss"

		b = append(b, bu)

		resp, err := txn.Mutate(context.Background(), &api.Mutation{SetJson: productJson})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Response: %s\n", resp.Json)
		con = con + 1
	}

	err = txn.Commit(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	return b

}

func (d *Db) GetBuyerById(buyer_id string) Buyer {
	variables := map[string]string{"$id1": buyer_id}
	const q = `query Me($id1: string)
		{
			me(func: eq(BuyerID, $id1)){
				BuyerID
				Name
				Age
				Date
			}
		}
	`
	//fmt.Println(variables)
	txn := d.NewReadOnlyTxn()
	resp, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n", resp.Json)

	type RootBuyer struct {
		Me []Buyer `json:"me"`
	}

	var b Buyer
	var r RootBuyer
	err = json.Unmarshal(resp.GetJson(), &r)
	fmt.Println(err)

	if err != nil {
		log.Fatal(err)
	}

	out, _ := json.MarshalIndent(r, "", "\t")
	fmt.Printf("%s\n", out)

	b.BuyerID = r.Me[0].BuyerID
	b.Name = r.Me[0].Name
	b.Age = r.Me[0].Age
	b.Date = r.Me[0].Date

	return b
}

func (d *Db) GetAllBuyers() []Buyer {
	const q = `
		{
			me(func: has(Age)){
				BuyerID
				Name
				Age
				Date
			}
		}
	`
	txn := d.NewReadOnlyTxn()
	resp, err := txn.Query(context.Background(), q)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n", resp.Json)

	type RootBuyers struct {
		Me []Buyer `json:"me"`
	}

	var b []Buyer
	var r RootBuyers
	err = json.Unmarshal(resp.GetJson(), &r)
	fmt.Println(err)

	if err != nil {
		log.Fatal(err)
	}

	out, _ := json.MarshalIndent(r, "", "\t")
	fmt.Printf("%s\n", out)

	b = r.Me

	return b
}
