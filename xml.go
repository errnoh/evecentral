package evecentral

// ------------------- quicklook ---------------------

type ECQuicklook struct {
	Version     string      `xml:"version,attr"`
	Method      string      `xml:"method,attr"`
	Item        int         `xml:"quicklook>item"`
	Itemname    string      `xml:"quicklook>itemname"`
	Regions     string      `xml:"quicklook>regions"`
	Hours       int         `xml:"quicklook>hours"`
	Minqty      int         `xml:"quicklook>minqty"`
	Sell_orders []*QLOrders `xml:"quicklook>sell_orders>order"`
	Buy_orders  []*QLOrders `xml:"quicklook>buy_orders>order"`
}

type QLOrders struct {
	Id            string  `xml:"id,attr"` // <order id="2456258421">
	Region        int     `xml:"region"`
	Station       int     `xml:"station"`
	Station_name  string  `xml:"station_name"`
	Security      float64 `xml:"security"`
	Range         int     `xml:"range"`
	Price         float64 `xml:"price"`
	Vol_remain    int     `xml:"vol_remain"`
	Min_volume    int     `xml:"min_volume"`
	Expires       string  `xml:"expires"`
	Reported_time string  `xml:"reported_time"`
}

// ------------------- marketstat ---------------------

type ECMarketstat struct {
	Version string    `xml:"version,attr"`
	Method  string    `xml:"method,attr"`
	Item    []*MSItem `xml:"marketstat>type"`
}

type MSItem struct {
	Id   int          `xml:"id,attr"`
	All  *MSPriceInfo `xml:"all"`
	Buy  *MSPriceInfo `xml:"buy"`
	Sell *MSPriceInfo `xml:"sell"`
}

type MSPriceInfo struct {
	Volume     float64 `xml:"volume"`
	Avg        float64 `xml:"avg"`
	Max        float64 `xml:"max"`
	Min        float64 `xml:"min"`
	Stddev     float64 `xml:"stddev"`
	Median     float64 `xml:"median"`
	Percentile float64 `xml:"percentile"`
}
