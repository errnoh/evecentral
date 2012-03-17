eve-central.com XML API bindings for Go
---------------------------------------

import "github.com/errnoh/evecentral"

http://eve-central.com/home/develop.html

#Usage example:

`    contents, ok := evecentral.Quicklook(34, -1, []int{}, -1, -1)`
`    contents, ok := evecentral.Marketstat(2, []int{34,35}, -1, []int{}, 30000142)`

_-1 is used as nil value for integer type parameters_

#TODO:
* Share a single decoder between API calls
