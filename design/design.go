package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var BasicAuth = BasicAuthSecurity("basic_auth")

var _ = API("lgbk", func() {
	Title("lgbk")
	Description("A journey's logbook named your life.")
	Scheme("http")
	Host("localhost:8082")
})

var _ = Resource("date", func() {
	BasePath("/dates")
	DefaultMedia(DateMedia)
	Security(BasicAuth)

	Action("index", func() {
		Description("Get all dates")
		Routing(GET("/"))
		Response(OK)
		Response(NotFound)
	})

	Action("show", func() {
		Description("Get date by id")
		Routing(GET("/:dateID"))
		Params(func() {
			Param("dateID", Integer, "Date ID")
		})
		Response(OK)
		Response(NotFound)
	})
})

// DateMedia defines
var DateMedia = MediaType("application/vnd.goa.example.date+json", func() {
	Description("A date in lgbk")
	Attributes(func() {
		Attribute("id", Integer, "Unique bottle ID")
		Attribute("title", String, "title of date")
		Attribute("body", String, "body of date")
		Required("id", "title", "body")
	})
	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("body")
	})
})
