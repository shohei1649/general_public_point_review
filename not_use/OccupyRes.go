package design

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
)

// OccupyRes dianping.order.occupyへのレスポンス結果
var OccupyRes = apidsl.MediaType("application/vnd.occupy_res+json", func() {

	apidsl.Attributes(func() {
		apidsl.Attribute("code", design.Integer, "Return code")
		apidsl.Attribute("isSuccess", design.Boolean, "Placeholder success")
		apidsl.Attribute("msg", design.String, "Tips")
		apidsl.Attribute("otaOrderId", design.String, "Suppliers order ID")
		apidsl.Attribute("orderId", design.Number, "New orders for big US ID")
		apidsl.Attribute("otaOrderStatus", design.Integer, "Order Status")

	})

	apidsl.View("default", func() {
		apidsl.Attribute("code")
		apidsl.Attribute("isSuccess")
		apidsl.Attribute("msg")
		apidsl.Attribute("otaOrderId")
		apidsl.Attribute("otaOrderStatus")
	})
})
