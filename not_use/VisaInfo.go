package design

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
)

// VisaInfo visa card Info
var VisaInfo = apidsl.Type("VisaInfo", func() {

	apidsl.Member("visaCountry", design.String, "National Visas", func() {
		apidsl.Default("visaCountry")
		apidsl.MaxLength(256)
		apidsl.Example("")
	})
	apidsl.Member("visaType", design.String, "Type of Visa", func() {
		apidsl.Default("visaType")
		apidsl.MaxLength(256)
		apidsl.Example("")
	})
})
