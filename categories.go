package main

import (
	"net/http"
	"encoding/json"
	"encoding/csv"
	"fmt"
	"strconv"
	"sync"
	"io/ioutil"
	"os"
)


func ScrapeFilmAndTv(fname string) {
	
	results, err := os.Create(fname)
	if err != nil {
		fmt.Println(err)
	}
	
	w := csv.NewWriter(results)
	w.Comma = '\t'
	defer w.Flush()
	
	var mu sync.Mutex
	var wg sync.WaitGroup
	

	w.Write(Header)
	wg.Add(1)
	go func() {
		firstRecord := 0

		for {
			data := ProductsCategories{}
			res, err := http.Get("https://wss2.cex.uk.webuy.io/v3/boxes?categoryIds=[932,681,793,40,747,850,792]&firstRecord="+strconv.Itoa(firstRecord)+"&count=50&sortBy=relevance&sortOrder=desc")
			if err != nil {
				fmt.Println(err)
			}
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
			}
			err = json.Unmarshal(body, &data)
			if data.Response.Data.Boxes == nil && data.Response.Ack == "Success" {
				fmt.Println("no more products", firstRecord, string(body))
				break
			}
			if data.Response.Data.Boxes == nil && data.Response.Ack == "Failure" {
				fmt.Println("Fail on record", firstRecord, string(body))
			}
			
			for _, v := range data.Response.Data.Boxes {
				var stock string
				if v.OutOfEcomStock == 0 {
					stock = "In Stock"
				} else {
					continue
				}
				mu.Lock()
				w.Write([]string{
					"0",
					v.BoxID,
					v.BoxName,
					fmt.Sprintf("%.2f", (v.SellPrice + 1.5)),
					stock,
				})
				mu.Unlock()

			}
			firstRecord += 50
		}
		wg.Done()
	
	}()
	wg.Add(1)
	go func() {
			firstRecord := 0
			
			for {
				data := ProductsCategories{}
				res, err := http.Get("https://wss2.cex.uk.webuy.io/v3/boxes?categoryIds=[746,749,710,708,843,991,845,709]&firstRecord="+strconv.Itoa(firstRecord)+"&count=50&sortBy=relevance&sortOrder=desc")
				if err != nil {
					fmt.Println(err)
				}
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					fmt.Println(err)
				}
				err = json.Unmarshal(body, &data)
				if data.Response.Data.Boxes == nil && data.Response.Ack == "Success" {
					fmt.Println("no more products", firstRecord, string(body))
					break
				}
				if data.Response.Data.Boxes == nil && data.Response.Ack == "Failure" {
					fmt.Println("Fail on record", firstRecord, string(body))
				}
				for _, v := range data.Response.Data.Boxes {
					var stock string
					if v.OutOfEcomStock == 0 {
						stock = "In Stock"
					} else {
						continue
					}
					
					mu.Lock()
					w.Write([]string{
						v.BoxID,
						v.BoxName,
						fmt.Sprintf("%.2f", v.SellPrice),
						stock,
					})
					mu.Unlock()
					
				}
				firstRecord += 50
			
			}
		wg.Done()	
		}()
	wg.Wait()
}
