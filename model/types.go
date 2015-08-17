package model

import (
	"time"

	"github.com/jmoiron/sqlx/types"
)

// AutoIncr contains fields that every db struct should have.
type AutoIncr struct {
	ID         int        `json:"id"`
	CreatedAt  time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time  `json:"updatedAt" db:"updated_at"`
	ArchivedAt *time.Time `json:"archivedAt" db:"archived_at"`
}

func (a *AutoIncr) UpdateTime() {
	if a.CreatedAt.IsZero() {
		a.CreatedAt = time.Now().UTC()
	}
	a.UpdatedAt = time.Now().UTC()

}

//Announcement ...
type Announcement struct {
	AutoIncr
	Title      string    `json:"title"`
	PersonID   int       `json:"personID" db:"person_id"` // Change to PostedBy?
	Author     *Person   `json:"author"`
	PostedDate time.Time `json:"postedDate" db:"posted_date"`
}

//Assignment ...
type Assignment struct {
	AutoIncr
	Name     string           `json:"name"`
	CourseID int              `json:"courseID" db:"course_id"`
	TermID   int              `json:"termID"`
	GroupID  int              `json:"groupID"`
	Group    *AssignmentGroup `json:"group"`
	MaxScore int16            `json:"maxScore"`
	DueDate  *time.Time       `json:"dueDate"`
}

type AssignmentAttempts struct {
	AutoIncr
	AssignmentID int     `json:"assignmentID"`
	PersonID     int     `json:"personID"`
	Score        string  `json:"score"`
	GradeAverage float32 `json:"gradeAverage"`
}

// AssignmentGroup ...
type AssignmentGroup struct {
	AutoIncr
	Name     string  `json:"name"`
	Weight   float64 `json:"weight"`
	CourseID int     `json:"courseID"`
	TermID   int     `json:"termID"`
}

// Course --
type Course struct {
	AutoIncr
	Name        string `json:"name"`
	LevelID     int    `json:"levelID"`
	GradeLevel  string `json:"gradeLevel"`
	MaxStudents int    `json:"maxStudents"`
	Terms       []Term `json:"terms"`
}

type CourseTerm struct {
	AutoIncr
	CourseID int
	TermID   int
}

// Enrollment --
type Enrollment struct {
	AutoIncr

	PersonID int             `json:"personID" db:"person_id"`
	Person   *types.JsonText `json:"person"`
	CourseID int             `json:"courseID" db:"course_id"`
	TermID   int             `json:"termID" db:"term_id"`
}

// Person --
type Person struct {
	AutoIncr
	FirstName   string         `json:"firstName" db:"first_name"`
	MiddleName  string         `json:"middleName" db:"middle_name"`
	LastName    string         `json:"lastName" db:"last_name"`
	Role        Role           `json:"role"`
	GradeLevel  string         `json:"gradeLevel" db:"grade_level"`
	ContactInfo *[]ContactInfo `json:"contactInfo"`
}

// ContactInfo --
type ContactInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}

// Session --
type Session struct {
	AutoIncr
	Token     string    `json:"token"`
	AccountID int       `json:"accountID"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// Term --
type Term struct {
	AutoIncr
	Name       string `json:"name"`
	SchoolYear int    `json:"schoolYear"`
}

type Level struct {
	AutoIncr
	Name string `json:"name"`
}

// Account --
type Account struct {
	AutoIncr
	PersonID        int    `json:"personID"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	HashedPassword  string `json:"-"`
	ActivationToken string `json:"-"` // base64 url encoded random hash.
	Disabled        bool   `json:"disabled"`
}
