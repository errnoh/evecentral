package evecentral

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

/*
<evec_api version...
    <quicklook>
        <sell_orders>
            order id=...>
                region>1</region>
               ....
           </order>
        </sell_orders>
        <buy_orders>
            ...
        </buy_orders>
    </quicklook>
</evec_api>
*/

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

func Marketstat(hours value, typeid list, minQ value, regionlimit list, usesystem value) (contents *MS, ok bool) {
        contents = new(MS)

	if len(typeid) == 0 {
		contents = nil
		return
	}

	resp, err := http.Get(fmt.Sprintf("http://api.eve-central.com/api/marketstat?%s%s%s%s%s", hours.insert("hours"), typeid.insert("typeid"), minQ.insert("minQ"), regionlimit.insert("regionlimit"), usesystem.insert("usesystem")))

	if err != nil {
		fmt.Println(err)
		contents = nil
		return
	}
	defer resp.Body.Close()

	//	body, err := ioutil.ReadAll(resp.Body)
	//	err = xml.Unmarshal(body, contents)
	// vs
	//      d := xml.NewDecoder(resp.Body)
	//      d.Decode(contents)

	d := xml.NewDecoder(resp.Body)
	d.Decode(contents)

	if err != nil {
		fmt.Println(err)
		contents = nil
		return
	}
	ok = true
	return
}

func Quicklook(typeid value, sethours value, regionlimit list, usesystem value, setminQ value) (contents *QL, ok bool) {
        contents = new(QL)

	if typeid <= 0 {
		contents = nil
		return
	}

	resp, err := http.Get(fmt.Sprintf("http://api.eve-central.com/api/quicklook?%s%s%s%s%s", typeid.insert("typeid"), sethours.insert("sethours"), regionlimit.insert("regionlimit"), usesystem.insert("usesystem"), setminQ.insert("setminQ")))

	if err != nil {
		fmt.Println(err)
		contents = nil
		return
	}
	defer resp.Body.Close()

	d := xml.NewDecoder(resp.Body)
	d.Decode(contents)

	if err != nil {
		fmt.Println(err)
		contents = nil
		return
	}
	ok = true
	return
}
