package model

import "time"

// AutoIncr contains fields that every db struct should have.
type AutoIncr struct {
	ID         int        `json:"id"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	ArchivedAt *time.Time `json:"archivedAt"`
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
	PersonID   int       `json:"personID"` // Change to PostedBy?
	Author     *Person   `json:"author"`
	PostedDate time.Time `json:"postedDate"`
}

//Assignment ...
type Assignment struct {
	AutoIncr
	Name     string           `json:"name"`
	CourseID int              `json:"courseID"`
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
	Name        string   `json:"name"`
	GradeLevel  string   `json:"gradeLevel"`
	MaxStudents int      `json:"maxStudents"`
	Terms       []string `json:"terms"`
}

// EmailConfirmation --
type EmailConfirmation struct {
	AutoIncr
	UserID int       `json:"userID"`
	UsedOn time.Time `json:"usedOn"`
}

// Enrollment --
type Enrollment struct {
	AutoIncr

	PersonID int     `json:"personID"`
	Person   *Person `json:"person"`
	CourseID int     `json:"courseID"`
	TermID   int     `json:"termID"`
}

// Person --
type Person struct {
	AutoIncr
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	// Types       []string `json:"types,omitempty"`
	GradeLevel  string         `json:"gradeLevel"`
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
	UserID    int       `json:"userID"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// Term --
type Term struct {
	AutoIncr
	Name       string    `json:"name"`
	SchoolYear time.Time `json:"schoolYear"`
}

// Type --
type Type struct{}

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
