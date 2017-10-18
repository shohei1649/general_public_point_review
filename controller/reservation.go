package controller

import (
	"github.com/goadesign/goa"
	"github.com/work/general_public_point_review/app"
)

// ReservationController implements the reservation resource.
type ReservationController struct {
	*goa.Controller
}

// NewReservationController creates a reservation controller.
func NewReservationController(service *goa.Service) *ReservationController {
	return &ReservationController{Controller: service.NewController("ReservationController")}
}

// Entry runs the entry action.
func (c *ReservationController) Entry(ctx *app.EntryReservationContext) error {
	// ReservationController_Entry: start_implement

	// Put your logic here

	// ReservationController_Entry: end_implement
	res := &app.GeneralPublicPointReview{
		Code:           200,
		IsSuccess:      true,
		Msg:            "OK",
		OrderID:        101101,
		OtaOrderID:     "10101",
		OtaOrderStatus: 1020,
	}
	return ctx.OK(res)
}
