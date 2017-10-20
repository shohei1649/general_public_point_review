// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "general_public_point_review": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/work/general_public_point_review/design
// --out=$(GOPATH)/src/github.com/work/general_public_point_review
// --version=v1.3.0

package app

import (
	"context"
	"net/http"

	"github.com/goadesign/goa"
)

// EntryReservationContext provides the reservation entry action context.
type EntryReservationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *OccupyPayLoad
}

// NewEntryReservationContext parses the incoming request URL and body, performs validations and creates the
// context used by the reservation controller entry action.
func NewEntryReservationContext(ctx context.Context, r *http.Request, service *goa.Service) (*EntryReservationContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := EntryReservationContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *EntryReservationContext) OK(r *GeneralPublicPointReview) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.general_public_point_review+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *EntryReservationContext) BadRequest(r *GeneralPublicPointReview) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.general_public_point_review+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *EntryReservationContext) Unauthorized(r *GeneralPublicPointReview) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.general_public_point_review+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *EntryReservationContext) InternalServerError(r *GeneralPublicPointReview) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.general_public_point_review+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}
