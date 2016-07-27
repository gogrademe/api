package main

import (
	"encoding/base64"

	"github.com/Sirupsen/logrus"
	h "github.com/gogrademe/api/handler"
	"github.com/gogrademe/api/store"
	"github.com/mattes/migrate/migrate"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	mw "github.com/labstack/echo/middleware"
	"github.com/mattaitchison/env"
)

var (
	listenAddr = env.String("listen_addr", ":5000", "listen address")
	dbAddr     = env.String("database_url", "postgres://localhost/gogrademe-api-dev?sslmode=disable&timezone=Etc/UTC", "sql db address")
	signingkey = env.String("jwt_key", "examplesigningkey", "key to used to sign jwt")
	// signingmethod = env.String("jwt_method", "HS256", "method used to sign jwt")
)

func main() {
	logrus.Println("running migrations")
	allErrors, ok := migrate.UpSync(dbAddr, "./migrations")
	if !ok {
		logrus.Println("Oh no ...")
		logrus.Fatal(allErrors)
		// do sth with allErrors slice
	}

	// Echo instance
	e := echo.New()
	e.Use(mw.LoggerWithConfig(mw.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(mw.Recover())
	e.Use(mw.CORS())

	s := store.Connect(dbAddr)
	s.EnsureAdmin()
	e.Use(h.SetDB(s))

	// catch sql no rows errors and return as a 404
	// e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		err := h(c)
	// 		if err == store.ErrNoRows {
	// 			return echo.NewHTTPError(http.StatusNotFound)
	// 		}
	// 		return err
	// 	}
	// })

	// e.Post("/session", h.CreateSession(signingkey, signingmethod))
	e.Post("/activate/:token", h.ActivateAccount)
	e.Get("/setup", h.CanSetup)
	e.Post("/setup", h.SetupApp)

	decoded, err := base64.URLEncoding.DecodeString(signingkey)
	if err != nil {
		panic(err)
	}

	auth := e.Group("")
	auth.Use(mw.JWT(decoded))
	auth.Get("/me", h.GetSession)
	// Accounts
	auth.Get("/account", h.GetAllAccounts)
	auth.Post("/account", h.CreateAccount)
	auth.Delete("/account/:id", h.DeleteAccount)

	// People
	g := auth.Group("/person")
	g.GET("", h.GetAllPeople)
	g.POST("", h.CreatePerson)
	g.GET("/:id", h.GetPerson)
	g.PUT("/:id", h.UpdatePerson)
	g.DELETE("/:id", h.DeletePerson)

	// Courses
	g = auth.Group("/course")
	g.GET("", h.GetAllCourses)
	g.POST("", h.CreateCourse)
	g.GET("/:id", h.GetCourse)
	g.PUT("/:id", h.UpdateCourse)
	g.PUT("/:courseID/term/:termID", h.CreateCourseTerm)
	g.DELETE("/:id", h.DeleteCourse)
	g.GET("/:courseID/term/:termID/assignments", h.GetCourseAssignments)
	g.GET("/:courseID/term/:termID/gradebook", h.GetGradebook)

	// Attempt
	g = auth.Group("/attempt")
	g.GET("", h.GetAllAttempts)
	g.POST("", h.CreateAttempt)
	g.GET("/:id", h.GetAttempt)
	g.PUT("/:id", h.UpdateAttempt)
	g.DELETE("/:id", h.DeleteAttempt)

	// Announcement
	g = auth.Group("/announcement")
	g.GET("", h.GetAllAnnouncements)
	g.POST("", h.CreateAnnouncement)
	g.GET("/:id", h.GetAnnouncement)
	g.PUT("/:id", h.UpdateAnnouncement)
	g.DELETE("/:id", h.DeleteAnnouncement)

	// Enrollments
	g = auth.Group("/enrollment")
	g.GET("", h.GetAllEnrollments)
	g.POST("", h.CreateEnrollment)
	g.GET("/:id", h.GetEnrollment)
	g.PUT("/:id", h.UpdateEnrollment)
	g.DELETE("/:id", h.DeleteEnrollment)

	// Terms
	g = auth.Group("/term")
	g.GET("", h.GetAllTerms)
	g.POST("", h.CreateTerm)
	g.GET("/:id", h.GetTerm)
	g.PUT("/:id", h.UpdateTerm)
	g.DELETE("/:id", h.DeleteTerm)

	// Levels
	g = auth.Group("/level")
	g.GET("", h.GetAllLevels)
	g.POST("", h.CreateLevel)
	g.GET("/:id", h.GetLevel)
	g.PUT("/:id", h.UpdateLevel)
	g.DELETE("/:id", h.DeleteLevel)

	// Assignments
	g = auth.Group("/assignment")
	g.GET("", h.GetAllAssignments)
	g.POST("", h.CreateAssignment)
	g.GET("/:id", h.GetAssignment)
	g.PUT("/:id", h.UpdateAssignment)
	g.DELETE("/:id", h.DeleteAssignment)

	// AssignmentGroups
	g = auth.Group("/group")
	g.GET("", h.GetAllAssignmentGroups)
	g.POST("", h.CreateAssignmentGroup)
	g.GET("/:id", h.GetAssignmentGroup)
	g.PUT("/:id", h.UpdateAssignmentGroup)
	g.DELETE("/:id", h.DeleteAssignmentGroup)

	// AssignmentGrades
	g = auth.Group("/grade")

	// Start server
	logrus.Println("Listening On:", listenAddr)

	e.Run(standard.New(listenAddr))
}
