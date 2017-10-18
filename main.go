//go:generate goagen bootstrap -d github.com/work/general_public_point_review/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/work/general_public_point_review/app"
	"github.com/work/general_public_point_review/controller"
)

func main() {
	// Create service
	service := goa.New("general_public_point_review")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "reservation" controller
	c := controller.NewReservationController(service)
	app.MountReservationController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
