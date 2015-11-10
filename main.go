package main

import (
	"github.com/Sirupsen/logrus"
	h "github.com/gogrademe/api/handler"
	"github.com/gogrademe/api/store"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/mattaitchison/env"
	"github.com/rs/cors"
)

var (
	listenAddr    = env.String("listen_addr", ":5000", "listen address")
	dbAddr        = env.String("db_addr", "postgres://localhost/gogrademe-api-dev?sslmode=disable&timezone=Etc/UTC", "sql db address")
	signingkey    = env.String("jwt_key", "examplesigningkey", "key to used to sign jwt")
	signingmethod = env.String("jwt_method", "HS256", "method used to sign jwt")
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Use(cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
	}).Handler)

	// Setup DB
	e.Use(h.SetDB(store.Connect(dbAddr)))

	e.Post("/session", h.CreateSession(signingkey, signingmethod))
	e.Post("/activate/:token", h.ActivateAccount)
	e.Get("/setup", h.CanSetup)
	e.Post("/setup", h.SetupApp)

	auth := e.Group("")
	auth.Use(h.JWTAuth(signingkey, signingmethod))

	// Accounts
	auth.Get("/account", h.GetAllAccounts)
	auth.Post("/account", h.CreateAccount)
	auth.Delete("/account/:id", h.DeleteAccount)

	// People
	g := auth.Group("/person")
	g.Get("", h.GetAllPeople)
	g.Post("", h.CreatePerson)
	g.Get("/:id", h.GetPerson)
	g.Put("/:id", h.UpdatePerson)
	g.Delete("/:id", h.DeletePerson)

	// Courses
	g = auth.Group("/course")
	g.Get("", h.GetAllCourses)
	g.Post("", h.CreateCourse)
	g.Get("/:id", h.GetCourse)
	g.Put("/:id", h.UpdateCourse)
	g.Put("/:courseID/term/:termID", h.CreateCourseTerm)
	g.Delete("/:id", h.DeleteCourse)
	g.Get("/:courseID/term/:termID/assignments", h.GetCourseAssignments)
	g.Get("/:courseID/term/:termID/gradebook", h.GetGradebook)

	// Attempt
	g = auth.Group("/attempt")
	g.Get("", h.GetAllAttempts)
	g.Post("", h.CreateAttempt)
	g.Get("/:id", h.GetAttempt)
	g.Put("/:id", h.UpdateAttempt)
	g.Delete("/:id", h.DeleteAttempt)

	// Announcement
	g = auth.Group("/announcement")
	g.Get("", h.GetAllAnnouncements)
	g.Post("", h.CreateAnnouncement)
	g.Get("/:id", h.GetAnnouncement)
	g.Put("/:id", h.UpdateAnnouncement)
	g.Delete("/:id", h.DeleteAnnouncement)

	// Enrollments
	g = auth.Group("/enrollment")
	g.Get("", h.GetAllEnrollments)
	g.Post("", h.CreateEnrollment)
	g.Get("/:id", h.GetEnrollment)
	g.Put("/:id", h.UpdateEnrollment)
	g.Delete("/:id", h.DeleteEnrollment)

	// Terms
	g = auth.Group("/term")
	g.Get("", h.GetAllTerms)
	g.Post("", h.CreateTerm)
	g.Get("/:id", h.GetTerm)
	g.Put("/:id", h.UpdateTerm)
	g.Delete("/:id", h.DeleteTerm)

	// Levels
	g = auth.Group("/level")
	g.Get("", h.GetAllLevels)
	g.Post("", h.CreateLevel)
	g.Get("/:id", h.GetLevel)
	g.Put("/:id", h.UpdateLevel)
	g.Delete("/:id", h.DeleteLevel)

	// Assignments
	g = auth.Group("/assignment")
	g.Get("", h.GetAllAssignments)
	g.Post("", h.CreateAssignment)
	g.Get("/:id", h.GetAssignment)
	g.Put("/:id", h.UpdateAssignment)
	g.Delete("/:id", h.DeleteAssignment)

	// AssignmentGroups
	g = auth.Group("/group")
	g.Get("", h.GetAllAssignmentGroups)
	g.Post("", h.CreateAssignmentGroup)
	g.Get("/:id", h.GetAssignmentGroup)
	g.Put("/:id", h.UpdateAssignmentGroup)
	g.Delete("/:id", h.DeleteAssignmentGroup)

	// AssignmentGrades
	g = auth.Group("/grade")

	// Start server
	logrus.Println("Listening On:", listenAddr)
	e.Run(listenAddr)
}
