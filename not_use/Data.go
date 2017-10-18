package design

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
)

// Data Data
var Data = apidsl.Type("Data", func() {
	apidsl.Member("orderId", design.Number, "New orders for big US ID", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("1000200030004000")
	})
	apidsl.Member("order price", design.Number, "Total order", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("20.0")
	})
	apidsl.Member("productId", design.Integer, "New American and product ID", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("12569")
	})
	apidsl.Member("otaPid", design.String, "Supplier Product ID", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("HB009")
	})
	apidsl.Member("otaPackageId", design.String, "Supplier Package ID", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("HBC001")
	})
	apidsl.Member("contact info", ContactInfo, "contact information", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("")
	})
	apidsl.Member("orderItems", OrderItems, "Orders Child", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("")
	})
	apidsl.Member("passengerInfos", PassengerInfos, "People travel information", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("")
	})
	apidsl.Member("wifiInfo", WifiInfo, "Wi-Fi 信息", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("")
	})
	apidsl.Member("flightInfo", design.String, "flight information", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("")
	})
	apidsl.Member("hotelInfo", design.String, "hotel info", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("")
	})
	apidsl.Member("Visinfao", VisaInfo, "Visa information", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("")
	})
	apidsl.Member("reserveConsumeTime", design.DateTime, "Food / recreational use of time", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("")
	})
	apidsl.Member("remark", design.String, "Remark", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("")
	})
	apidsl.Member("Cartype", design.String, "Models", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("")
	})

})
