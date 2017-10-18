package design

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
)

// PassengerInfo 予約者情報
var PassengerInfo = apidsl.Type("PassengerInfo", func() {

	apidsl.Member("saleType", design.Integer, "Type of sale", func() {
		apidsl.Minimum(0)
		apidsl.Maximum(128)
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("1")
	})
	apidsl.Member("name", design.String, "Play Name", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("New American and")
	})
	apidsl.Member("firstName", design.String, "Names play", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("College of Arts")
	})
	apidsl.Member("secondName", design.String, "Play man surnamed", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("new")
	})
	apidsl.Member("firstEnName", design.String, "Play Personal Name", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("meida")
	})
	apidsl.Member("secondEnName", design.String, "Play man surnamed Pinyin", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("please")
	})
	apidsl.Member("idCardNumber", design.String, "ID card", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("800000000")
	})
	apidsl.Member("passportNumber", design.String, "passport", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("P1234567")
	})
	apidsl.Member("hkPassportNumber", design.String, "Permit", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("P1234567")
	})
	apidsl.Member("twPassportNumber", design.String, "Taiwan Pass", func() {
		apidsl.Default("value")
		apidsl.MaxLength(256)
		apidsl.Example("T1234567")
	})

	apidsl.Required("saleType", "name", "firstName", "secondName", "idCardNumber", "passportNumber", "hkPassportNumber", "twPassportNumber")
})

//PassengerInfos List
var PassengerInfos = apidsl.ArrayOf(PassengerInfo)
