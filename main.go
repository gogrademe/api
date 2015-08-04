package main

import (
	"net/http"
	"net/url"

	h "github.com/gogrademe/api/handler"
	"github.com/jmoiron/sqlx"
	"github.com/serenize/snaker"

	"github.com/MattAitchison/envconfig"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

// const defaultDBName = "gogrademe-api-dev"
// const defaultDBDriver = "postgres"

var listenAddr = envconfig.String("listen_addr", ":5000", "listen address")
var dbAddr = envconfig.String("db_addr", "postgres://localhost/gogrademe-api-dev?sslmode=disable", "sql db address")

func connectDB(addr string) *sqlx.DB {
	dburi, _ := url.Parse(addr)
	db := sqlx.MustConnect(dburi.Scheme, dburi.String())
	db.MapperFunc(snaker.CamelToSnake)
	return db
}

// func bootstrapUser(db *sqlx.DB) error {
// 	as, err := store.GetAccountList(db)
// 	if err != nil {
// 		return err
// 	}
// 	if len(as) != 0 {
// 		return nil
// 	}
//
// 	a := model.NewAccountFor(1,"test@test.com")
// }

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Use(cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		Debug:          true,
	}).Handler)

	db := connectDB(dbAddr)

	// Setup DB
	e.Use(h.SetDB(db))
	e.Use(h.JWTAuth("someRandomSigningKey"))

	notmp := func(c *echo.Context) error {
		// d := h.ToDB(c)

		// e := d.DB().Ping()
		return c.JSON(http.StatusNotImplemented, e)
	}

	e.Post("/session", notmp)

	// auth := r.Group("", AuthRequired())

	// Accounts
	e.Get("/account", h.GetAllAccounts)
	e.Post("/account", h.CreateAccount)

	// People
	g := e.Group("/person")
	g.Get("", h.GetAllPeople)
	g.Post("", h.CreatePerson)
	g.Get("/:id", h.GetPerson)
	g.Put("/:id", notmp)
	g.Delete("/:id", notmp)

	// Courses
	g = e.Group("/course")
	g.Get("", notmp)
	g.Post("", notmp)
	g.Get("/:id", notmp)
	g.Put("/:id", notmp)
	g.Delete("/:id", notmp)

	// Enrollments
	g = e.Group("/enrollment")
	g.Get("", notmp)
	g.Post("", notmp)
	g.Get("/:id", notmp)
	g.Put("/:id", notmp)
	g.Delete("/:id", notmp)

	// Terms
	g = e.Group("/term")
	g.Get("", notmp)
	g.Post("", notmp)
	g.Get("/:id", notmp)
	g.Put("/:id", notmp)

	// SchoolYears
	g = e.Group("/schoolYear")
	g.Get("", notmp)
	g.Post("", notmp)
	g.Get("/:id", notmp)
	g.Put("/:id", notmp)

	// Assignments
	g = e.Group("/assignment")
	g.Get("", notmp)
	g.Post("", notmp)
	g.Get("/:id", notmp)
	g.Put("/:id", notmp)
	g.Delete("/:id", notmp)

	// AssignmentGroups
	g = e.Group("/assignmentGroup")
	g.Get("", notmp)
	g.Post("", notmp)
	g.Get("/:id", notmp)
	g.Put("/:id", notmp)
	g.Delete("/:id", notmp)

	// AssignmentGrades
	g = e.Group("/grade")
	g.Get("", notmp)
	g.Post("", notmp)
	// g.Get("/:id", , notmp)
	// g.Put("/:id", , notmp)

	// Start server
	e.Run(listenAddr)
}
