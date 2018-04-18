package main


type ProductsAll struct {
	Response struct {
		Ack  string `json:"ack"`
		Data struct {
			BoxlistsBoxes []struct {
				BoxID                string `json:"boxId"`
				BoxName              string `json:"boxName"`
				CategoryName         string `json:"categoryName"`
				CategoryFriendlyName string `json:"categoryFriendlyName"`
				SuperCatID           int    `json:"superCatId"`
				SuperCatName         string `json:"superCatName"`
				SuperCatFriendlyName string `json:"superCatFriendlyName"`
				ImageUrls            struct {
					Large  string `json:"large"`
					Medium string `json:"medium"`
					Small  string `json:"small"`
				} `json:"imageUrls"`
				IsNewBox       int         `json:"isNewBox"`
				SellPrice      float32         `json:"sellPrice"`
				CashPrice      int         `json:"cashPrice"`
				ExchangePrice  int         `json:"exchangePrice"`
				BoxRating      interface{} `json:"boxRating"`
				CategoryID     int         `json:"categoryId"`
				CannotBuy      int         `json:"cannotBuy"`
				OutOfEcomStock int         `json:"outOfEcomStock"`
			} `json:"boxlistsBoxes"`
		} `json:"data"`
		Error struct {
			Code            string        `json:"code"`
			InternalMessage string        `json:"internal_message"`
			MoreInfo        []interface{} `json:"moreInfo"`
		} `json:"error"`
	} `json:"response"`
}


var Header = []string{
	"Catno",
	"Barcode",
	"Title",
	"Price",
	"InStock",
}

type ProductsCategories struct {
	Response struct {
		Ack  string `json:"ack"`
		Data struct {
			Boxes []struct {
				BoxID                string `json:"boxId"`
				BoxName              string `json:"boxName"`
				IsMasterBox          int    `json:"isMasterBox"`
				CategoryID           int    `json:"categoryId"`
				CategoryName         string `json:"categoryName"`
				CategoryFriendlyName string `json:"categoryFriendlyName"`
				SuperCatID           int    `json:"superCatId"`
				SuperCatName         string `json:"superCatName"`
				SuperCatFriendlyName string `json:"superCatFriendlyName"`
				ImageUrls            struct {
					Large           string      `json:"large"`
					Medium          string      `json:"medium"`
					Small           string      `json:"small"`
					MasterBoxLarge  interface{} `json:"masterBoxLarge"`
					MasterBoxMedium interface{} `json:"masterBoxMedium"`
					MasterBoxSmall  interface{} `json:"masterBoxSmall"`
				} `json:"imageUrls"`
				CannotBuy      int `json:"cannotBuy"`
				IsNewBox       int `json:"isNewBox"`
				SellPrice      float32 `json:"sellPrice"`
				CashPrice      int `json:"cashPrice"`
				ExchangePrice  int `json:"exchangePrice"`
				BoxRating      int `json:"boxRating"`
				OutOfStock     int `json:"outOfStock"`
				OutOfEcomStock int `json:"outOfEcomStock"`
			} `json:"boxes"`
			TotalRecords int `json:"totalRecords"`
			MinPrice     int `json:"minPrice"`
			MaxPrice     int `json:"maxPrice"`
			Facets       struct {
				SuperCatName []struct {
					Name  string `json:"name"`
					ID    int    `json:"id"`
					Count int    `json:"count"`
				} `json:"superCatName"`
				CategoryFriendlyName []struct {
					Name  string `json:"name"`
					ID    int    `json:"id"`
					Count int    `json:"count"`
				} `json:"categoryFriendlyName"`
				ManufacturerName       interface{} `json:"manufacturerName"`
				NetworkName            interface{} `json:"networkName"`
				AttributeStructureInfo interface{} `json:"attributeStructureInfo"`
			} `json:"facets"`
		} `json:"data"`
		Error struct {
			Code            string        `json:"code"`
			InternalMessage string        `json:"internal_message"`
			MoreInfo        []interface{} `json:"moreInfo"`
		} `json:"error"`
	} `json:"response"`
}