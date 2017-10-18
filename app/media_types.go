// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "general_public_point_review": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/work/general_public_point_review/design
// --out=$(GOPATH)/src/github.com/work/general_public_point_review
// --version=v1.3.0

package app

import (
	"github.com/goadesign/goa"
)

// Occupy Response (default view)
//
// Identifier: application/vnd.general_public_point_review+json; view=default
type GeneralPublicPointReview struct {
	Code           int    `form:"code" json:"code" xml:"code"`
	IsSuccess      bool   `form:"isSuccess" json:"isSuccess" xml:"isSuccess"`
	Msg            string `form:"msg" json:"msg" xml:"msg"`
	OrderID        int    `form:"orderId,omitempty" json:"orderId,omitempty" xml:"orderId,omitempty"`
	OtaOrderID     string `form:"otaOrderId" json:"otaOrderId" xml:"otaOrderId"`
	OtaOrderStatus int    `form:"otaOrderStatus" json:"otaOrderStatus" xml:"otaOrderStatus"`
}

// Validate validates the GeneralPublicPointReview media type instance.
func (mt *GeneralPublicPointReview) Validate() (err error) {

	if mt.Msg == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "msg"))
	}
	if mt.OtaOrderID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "otaOrderId"))
	}

	return
}
