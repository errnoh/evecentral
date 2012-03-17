package evecentral

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"testing"
)

var testresponse, _ = http.Get("http://api.eve-central.com/api/marketstat?typeid=34&usesystem=30000142")

func TestMarketstat(t *testing.T) {
	_, err := Marketstat(2, []int{34, 35}, -1, []int{}, 30000142)
	if err != nil {
		t.Errorf("Failed to get Marketstat data: %s", err)
	}
}

func TestQuicklook(t *testing.T) {
	_, err := Quicklook(34, -1, []int{}, -1, -1)
	if err != nil {
		t.Errorf("Failed to get Quicklook data: %s", err)
	}
}

// Test functions for decode() solutions:
//
//	body, err := ioutil.ReadAll(resp.Body)
//	err = xml.Unmarshal(body, contents)
// vs
//      d := xml.NewDecoder(resp.Body)
//      err = d.Decode(contents)

func BenchmarkSharedDecoder(b *testing.B) {
	var contents *ECMarketstat
	d := xml.NewDecoder(testresponse.Body)

	for i := 0; i < b.N; i++ {
		contents = new(ECMarketstat)
		d.Decode(contents)
	}
	b.StopTimer()
	if contents == nil {
		b.Fatal("Decode failed, nil contents")
	}
}

func BenchmarkEachOwnDecoders(b *testing.B) {
	var contents *ECMarketstat

	for i := 0; i < b.N; i++ {
		d := xml.NewDecoder(testresponse.Body)
		contents = new(ECMarketstat)
		d.Decode(contents)
	}
	b.StopTimer()
	if contents == nil {
		b.Fatal("Decode failed, nil contents")
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	var contents *ECMarketstat

	for i := 0; i < b.N; i++ {
		contents = new(ECMarketstat)
		body, _ := ioutil.ReadAll(testresponse.Body)
		xml.Unmarshal(body, contents)
	}
	b.StopTimer()
	if contents == nil {
		b.Fatal("Unmarshal failed, nil contents")
	}
}
