package store

import "github.com/gogrademe/api/model"

// GetAnnouncement --
func (s *Store) GetAnnouncement(id int) (*model.Announcement, error) {
	var r model.Announcement
	return &r, s.db.Get(&r, "select * from announcement WHERE id=$1", id)
}

// GetAnnouncementList --
func (s *Store) GetAnnouncementList() ([]model.Announcement, error) {
	var r []model.Announcement
	return r, s.db.Select(&r, "select * from announcement")
}

// InsertAnnouncement --
func (s *Store) InsertAnnouncement(announcement *model.Announcement) error {
	stmt := `INSERT INTO announcement (title, body, posted_date, person_id, created_at, updated_at)
			 VALUES (:title, :body, :posted_date, :person_id, :created_at, :updated_at) RETURNING id`
	announcement.UpdateTime()

	var err error
	announcement.ID, err = insert(s.db, stmt, announcement)
	return err
}

// Update --
func (s *Store) UpdateAnnouncement(announcement *model.Announcement) error {
	stmt := Update("announcement").SetN("title", "body", "posted_date", "person_id", "created_at", "updated_at").Eq("id").String()
	announcement.UpdateTime()

	_, err := s.db.NamedQuery(stmt, announcement)
	return err

}

// Del --
func (s *Store) DeleteAnnouncement(id int) error {
	stmt := `DELETE FROM announcement WHERE id=$1`

	_, err := s.db.Exec(stmt, id)
	return err

}
