// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "general_public_point_review": reservation TestHelpers
//
// Command:
// $ goagen
// --design=github.com/work/general_public_point_review/design
// --out=$(GOPATH)/src/github.com/work/general_public_point_review
// --version=v1.3.0

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/work/general_public_point_review/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// EntryReservationBadRequest runs the method Entry of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func EntryReservationBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ReservationController, payload *app.OccupyPayLoad) (http.ResponseWriter, *app.GeneralPublicPointReview) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/general_public_point_review/reservation"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ReservationTest"), rw, req, prms)
	entryCtx, __err := app.NewEntryReservationContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	entryCtx.Payload = payload

	// Perform action
	__err = ctrl.Entry(entryCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt *app.GeneralPublicPointReview
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.GeneralPublicPointReview)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.GeneralPublicPointReview", resp, resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// EntryReservationInternalServerError runs the method Entry of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func EntryReservationInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ReservationController, payload *app.OccupyPayLoad) (http.ResponseWriter, *app.GeneralPublicPointReview) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/general_public_point_review/reservation"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ReservationTest"), rw, req, prms)
	entryCtx, __err := app.NewEntryReservationContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	entryCtx.Payload = payload

	// Perform action
	__err = ctrl.Entry(entryCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt *app.GeneralPublicPointReview
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.GeneralPublicPointReview)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.GeneralPublicPointReview", resp, resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// EntryReservationOK runs the method Entry of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func EntryReservationOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ReservationController, payload *app.OccupyPayLoad) (http.ResponseWriter, *app.GeneralPublicPointReview) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/general_public_point_review/reservation"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ReservationTest"), rw, req, prms)
	entryCtx, __err := app.NewEntryReservationContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	entryCtx.Payload = payload

	// Perform action
	__err = ctrl.Entry(entryCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.GeneralPublicPointReview
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.GeneralPublicPointReview)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.GeneralPublicPointReview", resp, resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// EntryReservationUnauthorized runs the method Entry of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func EntryReservationUnauthorized(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ReservationController, payload *app.OccupyPayLoad) (http.ResponseWriter, *app.GeneralPublicPointReview) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/general_public_point_review/reservation"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ReservationTest"), rw, req, prms)
	entryCtx, __err := app.NewEntryReservationContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	entryCtx.Payload = payload

	// Perform action
	__err = ctrl.Entry(entryCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 401 {
		t.Errorf("invalid response status code: got %+v, expected 401", rw.Code)
	}
	var mt *app.GeneralPublicPointReview
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.GeneralPublicPointReview)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.GeneralPublicPointReview", resp, resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}
