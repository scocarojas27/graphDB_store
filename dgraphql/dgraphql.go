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
	TransactionID string   `json:"TransactionID,omitempty"`
	BuyerID       string   `json:"BuyerID,omitempty"`
	Ip            string   `json:"Ip,omitempty"`
	Device        string   `json:"Device,omitempty"`
	Products      []string `json:"Products,omitempty"`
	Date          string   `json:"Date,omitempty"`
}

type Report struct {
	Transactions   []Transaction `json:"Transactions,omitempty"`
	SameIp         []Buyer       `json:"SameIp,omitempty"`
	Recomendations []Product     `json:"Recomendations,omitempty"`
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
		fmt.Printf("%s\n", resp.Json)
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

	type RootInsertBuyers struct {
		Me []Buyer `json:"me"`
	}

	var buyers []Buyer

	var data []interface{}
	json.Unmarshal(responseData, &data)
	insertDate := strconv.FormatInt(time.Now().Unix(), 10)
	con := 0

	for _, e := range data {
		var r RootInsertBuyers
		var bu Buyer
		uid := "_:Buyer" + strconv.Itoa(con) + insertDate
		b := e.(map[string]interface{})
		age := int(b["age"].(float64))
		buyer := map[string]interface{}{
			"uid":     uid,
			"BuyerID": b["id"].(string),
			"Name":    b["name"].(string),
			"Age":     age,
			"Date":    insertDate,
		}
		productJson, err := json.Marshal(buyer)

		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(productJson, &r)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(err)

		bu.BuyerID = "ssss"
		bu.Name = "sssss"
		bu.Age = 1400
		bu.Date = "sssss"

		buyers = append(buyers, bu)

		resp, err := txn.Mutate(context.Background(), &api.Mutation{SetJson: productJson})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", resp.Json)
		con = con + 1
	}

	err = txn.Commit(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	return buyers
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
	//fmt.Printf("Response: %s\n", resp.Json)

	type RootBuyer struct {
		Me []Buyer `json:"me"`
	}

	var b Buyer
	var r RootBuyer
	err = json.Unmarshal(resp.GetJson(), &r)
	//fmt.Println(err)

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

func (d *Db) GetProductById(product_id string) (Product, error) {
	variables := map[string]string{"$id1": product_id}
	var p Product
	const q = `query Me($id1: string)
		{
			me(func: eq(ProductID, $id1)) @filter(has(Price)){
				ProductID
				Name
				Price
				Date
			}
		}
	`
	//fmt.Println(variables)
	txn := d.NewReadOnlyTxn()
	resp, err := txn.QueryWithVars(context.Background(), q, variables)
	fmt.Println("El puto json: ", &resp.Json)
	fmt.Println("Este es el hp error: ", err)
	if err != nil {
		log.Fatal(err)
	}

	type RootProduct struct {
		Me []Product `json:"me"`
	}

	var r RootProduct
	err = json.Unmarshal(resp.GetJson(), &r)

	if err != nil || len(r.Me) == 0 {
		return p, err
	}

	out, _ := json.MarshalIndent(r, "", "\t")
	fmt.Printf("%s\n", out)

	p.ProductID = r.Me[0].ProductID
	p.Name = r.Me[0].Name
	p.Price = r.Me[0].Price
	p.Date = r.Me[0].Date

	return p, nil
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
	//fmt.Printf("Response: %s\n", resp.Json)

	type RootBuyers struct {
		Me []Buyer `json:"me"`
	}

	var b []Buyer
	var r RootBuyers
	err = json.Unmarshal(resp.GetJson(), &r)
	//fmt.Println(err)

	if err != nil {
		log.Fatal(err)
	}

	out, _ := json.MarshalIndent(r, "", "\t")
	fmt.Printf("%s\n", out)

	b = r.Me

	return b
}

func (d *Db) GetTransactionsByBuyerId(buyer_id string) []Transaction {

	variables := map[string]string{"$id1": buyer_id}
	const q = `query Me($id1: string)
		{
			me(func: eq(BuyerID, $id1)) @filter(has(Ip)){
				TransactionID
				BuyerID
				Ip
				Device
				Products
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
	//fmt.Printf("Response: %s\n", resp.Json)

	type RootTransactions struct {
		Me []Transaction `json:"me"`
	}

	var t []Transaction
	var r RootTransactions
	err = json.Unmarshal(resp.GetJson(), &r)
	//fmt.Println(err)

	if err != nil {
		log.Fatal(err)
	}

	//out, _ := json.MarshalIndent(r, "", "\t")
	//fmt.Printf("%s\n", out)

	t = r.Me

	return t
}

func (d *Db) GetBuyersByIp(ip string) []Buyer {
	var buyers []Buyer
	variables := map[string]string{"$ip1": ip}
	const q1 = `query Me($ip1: string)
		{
			me(func: eq(Ip, $ip1)){
				TransactionID
				BuyerID
				Ip
				Device
				Products
				Date
			}
		}
	`
	txn := d.NewReadOnlyTxn()
	resp, err := txn.QueryWithVars(context.Background(), q1, variables)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Response: %s\n", resp.Json)

	type RootTransactions struct {
		Me []Transaction `json:"me"`
	}

	var tb []Transaction
	var r RootTransactions
	err = json.Unmarshal(resp.GetJson(), &r)
	//fmt.Println(err)

	if err != nil {
		log.Fatal(err)
	}

	//out, _ := json.MarshalIndent(r, "", "\t")
	//fmt.Printf("%s\n", out)

	tb = r.Me

	for _, t := range tb {
		buyer := d.GetBuyerById(t.BuyerID)
		if !(buyerInSlice(buyer, buyers)) {
			buyers = append(buyers, buyer)
		} else {
			fmt.Println("El comprador ya está en el listado")
		}
	}

	return buyers
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func buyerInSlice(a Buyer, list []Buyer) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (d *Db) GetReport(buyer_id string) Report {
	var report Report
	var products []string
	var recomedations []Product
	transactions := d.GetTransactionsByBuyerId(buyer_id)
	sameIp := d.GetBuyersByIp(transactions[0].Ip)

	for _, b := range sameIp {
		t := d.GetTransactionsByBuyerId(b.BuyerID)
		for _, p := range t {
			if !(stringInSlice(p.Products[0], products)) {
				products = append(products, p.Products[0])
			} else {
				fmt.Println("El producto ya está en el listado de recomendados.")
			}
		}
	}

	for _, r := range products {
		fmt.Println(r)
		p, err := d.GetProductById(r)
		if err != nil {
			log.Fatal(err)
		}
		if p.Name != "" {
			recomedations = append(recomedations, p)
		} else {
			fmt.Println("El producto no existe.")
		}

	}

	report.Transactions = transactions
	report.SameIp = sameIp
	report.Recomendations = recomedations

	return report
}
