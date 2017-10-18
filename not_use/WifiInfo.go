package design

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
)

// WifiInfo wifi info
var WifiInfo = apidsl.Type("WifiInfo", func() {
	apidsl.Member("getAddress", design.String, "Wi-Fi 取 件 地址", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("")
	})
	apidsl.Member("returnAddress", design.String, "Wi-Fi 还 件 地址", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("")
	})
	apidsl.Member("wifiCityName", design.String, "Wi-Fi 城市", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("")
	})

})
