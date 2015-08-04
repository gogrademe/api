-- Assignments
ALTER TABLE assignment_attempt ADD FOREIGN KEY (assignment_id) REFERENCES assignment (id);

-- AssignmentGroups
ALTER TABLE assignment ADD FOREIGN KEY (group_id) REFERENCES assignment_group (id);

-- Courses
ALTER TABLE assignment ADD FOREIGN KEY (course_id) REFERENCES course (id);
ALTER TABLE enrollment ADD FOREIGN KEY (course_id) REFERENCES course (id);
ALTER TABLE assignment_group ADD FOREIGN KEY (course_id) REFERENCES course (id);

-- People
ALTER TABLE announcement ADD FOREIGN KEY (person_id) REFERENCES person (id);
ALTER TABLE assignment_attempt ADD FOREIGN KEY (person_id) REFERENCES person (id);
ALTER TABLE enrollment ADD FOREIGN KEY (person_id) REFERENCES person (id);
ALTER TABLE account ADD FOREIGN KEY (person_id) REFERENCES person (id);

-- Terms
ALTER TABLE assignment ADD FOREIGN KEY (term_id) REFERENCES term (id);
ALTER TABLE assignment_group ADD FOREIGN KEY (term_id) REFERENCES term (id);
ALTER TABLE enrollment ADD FOREIGN KEY (term_id) REFERENCES term (id);

-- Users
ALTER TABLE email_confirmation ADD FOREIGN KEY (account_id) REFERENCES account (id);
ALTER TABLE session ADD FOREIGN KEY (account_id) REFERENCES account (id);

ALTER TABLE course ADD FOREIGN KEY (level_id) REFERENCES level (id);
