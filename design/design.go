package design // The convention consists of naming the design
// package "design"
import (
	"github.com/goadesign/goa/design" // Use . imports to enable the DSL
	"github.com/goadesign/goa/design/apidsl"
)

var _ = apidsl.API("general_public_point_review", func() { // API defines the microservice endpoint and

	apidsl.Title("general_public_point_review API")
	apidsl.Description("general_public_point_review_API")
	apidsl.Version("1.0.0")
	apidsl.Scheme("http")
	apidsl.Consumes("application/json")
	apidsl.Produces("application/json")
	apidsl.Host("localhost:8080")
})

var _ = apidsl.Resource("reservation", func() { // Resources group related API endpoints
	apidsl.BasePath("/general_public_point_review") // together. They map to REST resources for REST
	apidsl.Action("entry", func() {
		apidsl.Routing(apidsl.POST("/reservation"))
		apidsl.Payload(OccupyPayload)
		apidsl.Response(design.OK, OccupyResposne)
		apidsl.Response(design.BadRequest, OccupyResposne)
		apidsl.Response(design.Unauthorized, OccupyResposne)
		apidsl.Response(design.InternalServerError, OccupyResposne)

	})
})

var OccupyPayload = apidsl.Type("OccupyPayLoad", func() {
	apidsl.Member("otaId", design.Integer)
	apidsl.Member("data", design.String)
	apidsl.Member("sign", design.String)
	apidsl.Member("agentId", design.Integer)
	apidsl.Required("otaId", "data", "sign")

})

var OccupyResposne = apidsl.MediaType("application/vnd.general_public_point_review+json", func() {

	apidsl.Description("Occupy Response")
	apidsl.Attributes(func() {
		apidsl.Attribute("code", design.Integer)
		apidsl.Attribute("isSuccess", design.Boolean)
		apidsl.Attribute("msg", design.String)
		apidsl.Attribute("otaOrderId", design.String)
		apidsl.Attribute("orderId", design.Integer)
		apidsl.Attribute("otaOrderStatus", design.Integer)
		apidsl.Required("code", "isSuccess", "msg", "otaOrderId", "otaOrderStatus")
	})
	apidsl.View("default", func() {
		apidsl.Attribute("code")
		apidsl.Attribute("isSuccess")
		apidsl.Attribute("msg")
		apidsl.Attribute("otaOrderId")
		apidsl.Attribute("orderId")
		apidsl.Attribute("otaOrderStatus")
	})
})
