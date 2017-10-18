package design

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
)

// ContactInfo 連絡先情報
var ContactInfo = apidsl.Type("ContactInfo", func() {
	apidsl.Member("name", design.String, "Contact name", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("New American and")
	})
	apidsl.Member("phone", design.String, "contact number", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("13900000000")
	})
	apidsl.Member("email", design.String, "E-mail address", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("xinmeida@meituan.com")
	})
	apidsl.Member("pickupAddress", design.String, "Pick-up address", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("New American and")
	})
	apidsl.Member("hotelAddressEn", design.String, "Hotel address", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("Shangri-La's North Gate")
	})
	apidsl.Member("mail address", design.String, "Mailing address", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("Shanghai New American and")
	})
	apidsl.Member("mailZipCode", design.String, "Zip Code", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("210000")
	})
	apidsl.Member("startDate", design.DateTime, "Departure Date", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("2017-01-16 00:00:00")
	})
	apidsl.Member("endDate", design.DateTime, "End Date", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("2017-01-16 00:00:00")
	})

	apidsl.Required("name",
		"phone",
		"email",
		"pickupAddress",
		"hotelAddressEn",
		"mailAddress",
		"startDate") // ← 必須要素の指定
})
