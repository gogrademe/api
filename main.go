package main

import (
	"net/url"

	h "github.com/gogrademe/api/handler"
	m "github.com/gogrademe/api/model"

	"github.com/MattAitchison/envconfig"
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	_ "github.com/lib/pq"
)

// const defaultDBName = "gogrademe-api-dev"
// const defaultDBDriver = "postgres"

var listenAddr = envconfig.String("listen_addr", ":5000", "listen address")
var dbAddr = envconfig.String("db_addr", "postgres://localhost/gogrademe-api-dev?sslmode=disable", "sql db address")

func mustConnectDB(addr string) *gorm.DB {
	dburi, _ := url.Parse(addr)
	db, err := gorm.Open(dburi.Scheme, dburi.String())
	if err != nil {
		logrus.Fatal(err)
	}

	if err := db.DB().Ping(); err != nil {
		logrus.Fatal(err)
	}

	return &db
}

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	db := mustConnectDB(dbAddr)
	db.AutoMigrate(&m.User{}, &m.Person{}, &m.Assignment{}, &m.Session{}, &m.Announcement{})

	// Setup DB
	e.Use(h.SetDB(db))

	notmp := func(c *echo.Context) error {
		d := h.ToDB(c)

		e := d.DB().Ping()
		return c.JSON(500, e)
	}

	e.Post("/session", notmp)

	// auth := r.Group("", AuthRequired())

	// Users
	e.Get("/user", h.GetAllUsers)
	e.Post("/user", h.CreateUser)

	// People
	g := e.Group("/person")
	g.Get("", notmp)
	g.Post("", notmp)
	g.Get("/:id", notmp)
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
