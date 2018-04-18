package main

import (
	"net/http"
	"encoding/json"
	"encoding/csv"
	"fmt"
	"strconv"
	"io/ioutil"
	"os"
)


func ScrapeAll(fname string) {
	firstRecord := 0
	hasData := true
	
	results, err := os.Create(fname)
	if err != nil {
		fmt.Println(err)
	}
	
	w := csv.NewWriter(results)
	defer w.Flush()
	
	w.Write(Header)
	
	for hasData {
		data := ProductsAll{}
		res, err := http.Get("https://wss2.cex.uk.webuy.io/v3/boxlists/hotproducts?superCatId=*&firstRecord="+strconv.Itoa(firstRecord)+"&count=50")
		if err != nil {
			fmt.Println(err)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(body, &data)
		if data.Response.Data.BoxlistsBoxes == nil && data.Response.Ack == "Success" {
			fmt.Println("no more products", firstRecord, string(body))
			break
		}
		if data.Response.Data.BoxlistsBoxes == nil && data.Response.Ack == "Failure" {
			fmt.Println("Fail on record", firstRecord, string(body))
		}
		for _, v := range data.Response.Data.BoxlistsBoxes {
			var stock string
			if v.OutOfEcomStock == 0 {
				stock = "In Stock"
			} else {
				stock = "Out of Stock"
			}
			w.Write([]string{
				v.BoxID,
				v.BoxName,
				fmt.Sprintf("%.2f", v.SellPrice),
				stock,
			})
		}
		firstRecord += 50
	}
}