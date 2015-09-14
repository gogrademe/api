package model

import "time"

// AutoIncr contains fields that every db struct should have.
type AutoIncr struct {
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	ArchivedAt *time.Time `json:"archived_at"`
}

//Announcement ...
type Announcement struct {
	AutoIncr
	AnnouncementID int        `json:"announcement_id"`
	Title          string     `json:"title"`
	PersonID       int        `json:"person_id"` // Change to PostedBy?
	Author         *Person    `json:"author"`
	PostedDate     *time.Time `json:"posted_date"`
}

//Assignment ...
type Assignment struct {
	AutoIncr
	AssignmentID int              `json:"assignment_id"`
	Name         string           `json:"name"`
	CourseID     int              `json:"course_id"`
	TermID       int              `json:"term_id"`
	GroupID      int              `json:"group_id"`
	Group        *AssignmentGroup `json:"group"`
	MaxScore     int16            `json:"max_score"`
	DueDate      *time.Time       `json:"due_date"`
}

// Attempt represents a score on an assignment
type Attempt struct {
	AutoIncr
	AttemptID    int     `json:"attempt_id"`
	AssignmentID int     `json:"assignment_id"`
	PersonID     int     `json:"person_id"`
	Score        string  `json:"score"`
	Average      float32 `json:"grade_average"`
}

// AssignmentGroup ...
type AssignmentGroup struct {
	AutoIncr
	GroupID  int     `json:"group_id"`
	Name     string  `json:"name"`
	Weight   float64 `json:"weight"`
	CourseID int     `json:"course_id"`
	TermID   int     `json:"term_id"`
}

// Course --
type Course struct {
	AutoIncr
	CourseID    int    `json:"course_id"`
	Name        string `json:"name"`
	LevelID     int    `json:"level_id"`
	GradeLevel  string `json:"grade_level"`
	MaxStudents int    `json:"max_students"`
	Terms       []Term `json:"terms"`
}

type CourseTerm struct {
	AutoIncr
	CourseID int `json:"course_id"`
	TermID   int `json:"term_id"`
}

// Enrollment --
type Enrollment struct {
	AutoIncr
	EnrollmentID int     `json:"enrollment_id"`
	PersonID     int     `json:"person_id"`
	Person       *Person `json:"person"`
	CourseID     int     `json:"course_id"`
	TermID       int     `json:"term_id"`
}

// Person --
type Person struct {
	AutoIncr
	PersonID    int            `json:"person_id"`
	FirstName   string         `json:"first_name"`
	MiddleName  string         `json:"middle_name"`
	LastName    string         `json:"last_name"`
	Role        Role           `json:"role"`
	GradeLevel  string         `json:"grade_level"`
	ContactInfo *[]ContactInfo `json:"contact_info"`
}

// ContactInfo --
type ContactInfo struct {
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

// Session --
type Session struct {
	AutoIncr
	SessionID int       `json:"session_id"`
	Token     string    `json:"token"`
	AccountID int       `json:"account_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Term --
type Term struct {
	AutoIncr
	TermID     int    `json:"term_id"`
	Name       string `json:"name"`
	SchoolYear int    `json:"school_year"`
}

// Level represents a grade level
type Level struct {
	AutoIncr
	LevelID int    `json:"level_id"`
	Name    string `json:"name"`
}

// Account --
type Account struct {
	AutoIncr
	AccountID       int    `json:"account_id"`
	PersonID        int    `json:"person_id"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	HashedPassword  string `json:"-"`
	ActivationToken string `json:"-"` // base64 url encoded random hash.
	Disabled        bool   `json:"disabled"`
}
