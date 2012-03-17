package evecentral

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type XmlError string

func (e XmlError) Error() string {
	return string(e)
}

type list []int

func (l list) insert(name string) string {
	var s string
	for _, v := range l {
		s += fmt.Sprintf("&%s=%d", name, v)
	}
	return s
}

type value int

func (v value) insert(name string) string {
	if v == -1 {
		return ""
	}
	return fmt.Sprintf("&%s=%d", name, v)
}

func Marketstat(hours value, typeid list, minQ value, regionlimit list, usesystem value) (contents *ECMarketstat, err error) {
	contents = new(ECMarketstat)

	if len(typeid) <= 0 {
		contents = nil
		err = XmlError("Invalid typeid")
		return
	}

	err = decode(contents, fmt.Sprintf("http://api.eve-central.com/api/marketstat?%s%s%s%s%s", hours.insert("hours"), typeid.insert("typeid"), minQ.insert("minQ"), regionlimit.insert("regionlimit"), usesystem.insert("usesystem")))

	return
}

func Quicklook(typeid value, sethours value, regionlimit list, usesystem value, setminQ value) (contents *ECQuicklook, err error) {
	contents = new(ECQuicklook)

	if typeid <= 0 {
		contents = nil
		err = XmlError("Invalid typeid")
		return
	}

	err = decode(contents, fmt.Sprintf("http://api.eve-central.com/api/quicklook?%s%s%s%s%s", typeid.insert("typeid"), sethours.insert("sethours"), regionlimit.insert("regionlimit"), usesystem.insert("usesystem"), setminQ.insert("setminQ")))

	return
}

func decode(contents interface{}, url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	d := xml.NewDecoder(resp.Body)
	err = d.Decode(contents)

	if err != nil {
		return err
	}

	return err
}
