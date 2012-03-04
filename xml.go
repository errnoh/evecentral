package evecentral

// ------------------- quicklook ---------------------

type QL struct {
	Version   string       `xml:"version,attr"`
	Method    string       `xml:"method,attr"`
	Quicklook *QLQuicklook `xml:"quicklook"`
}

type QLQuicklook struct {
	Item        int       `xml:"item"`
	Itemname    string    `xml:"itemname"`
	Regions     string    `xml:"regions"`
	Hours       int       `xml:"hours"`
	Minqty      int       `xml:"minqty"`
	Sell_orders *QLOrders `xml:"sell_orders"`
	Buy_orders  *QLOrders `xml:"buy_orders"`
}

type QLOrders struct {
	Orders []*QLOrder `xml:"order"`
}

type QLOrder struct {
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

type MS struct {
	Version    string        `xml:"version,attr"`
	Method     string        `xml:"method,attr"`
	Marketstat *MSMarketstat `xml:"marketstat"`
}

type MSMarketstat struct {
	Types []*MSType `xml:"type"`
}

type MSType struct {
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
