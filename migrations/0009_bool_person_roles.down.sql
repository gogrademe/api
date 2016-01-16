ALTER TABLE person
    DROP COLUMN is_admin,
    DROP COLUMN is_teacher,
    DROP COLUMN is_parent,
    DROP COLUMN is_student,
    ADD COLUMN role integer NOT NULL;
