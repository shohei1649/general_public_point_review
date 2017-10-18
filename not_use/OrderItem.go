package design

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
)

// OrderItem Order Info
var OrderItem = apidsl.Type("OrderItem", func() {

	apidsl.Member("orderId", design.Number, "New orders for big US ID", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("1000200030004000")
	})
	apidsl.Member("skuId", design.Integer, "New American and SKUID", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("10065")
	})
	apidsl.Member("otaSkuId", design.String, "Supplier SKUID", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("RS0001")
	})
	apidsl.Member("saleType", design.String, "Type of sale", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("1")
	})
	apidsl.Member("quantity", design.String, "quantity", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("1")
	})
	apidsl.Member("skuPrice", design.String, "sku 单价", func() {
		apidsl.Default("value")
		apidsl.Enum(1, 2, 3)
		apidsl.MaxLength(256)
		apidsl.Example("10.0")
	})

})

//OrderItems OrderList
var OrderItems = apidsl.ArrayOf(OrderItem)
