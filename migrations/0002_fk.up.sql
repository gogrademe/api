-- Assignments
ALTER TABLE attempt ADD FOREIGN KEY (assignment_id) REFERENCES assignment (assignment_id);

-- AssignmentGroups
ALTER TABLE assignment ADD FOREIGN KEY (assignment_group_id) REFERENCES assignment_group (assignment_group_id);

-- Courses
ALTER TABLE assignment ADD FOREIGN KEY (course_id) REFERENCES course (course_id);
ALTER TABLE enrollment ADD FOREIGN KEY (course_id) REFERENCES course (course_id);
ALTER TABLE course_term ADD FOREIGN KEY (course_id) REFERENCES course (course_id);
-- People
ALTER TABLE announcement ADD FOREIGN KEY (person_id) REFERENCES person (person_id);
ALTER TABLE attempt ADD FOREIGN KEY (person_id) REFERENCES person (person_id);
ALTER TABLE enrollment ADD FOREIGN KEY (person_id) REFERENCES person (person_id);
ALTER TABLE account ADD FOREIGN KEY (person_id) REFERENCES person (person_id);

-- Terms
ALTER TABLE assignment ADD FOREIGN KEY (term_id) REFERENCES term (term_id);

ALTER TABLE enrollment ADD FOREIGN KEY (term_id) REFERENCES term (term_id);
ALTER TABLE course_term ADD FOREIGN KEY (term_id) REFERENCES term (term_id);

-- Users
ALTER TABLE email_confirmation ADD FOREIGN KEY (account_id) REFERENCES account (account_id);
ALTER TABLE session ADD FOREIGN KEY (account_id) REFERENCES account (account_id);

-- Level
ALTER TABLE course ADD FOREIGN KEY (level_id) REFERENCES level (level_id);

-- CourseTerm
ALTER TABLE assignment_group ADD FOREIGN KEY (course_id, term_id) REFERENCES course_term (course_id, term_id);
