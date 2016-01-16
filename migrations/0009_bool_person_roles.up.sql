ALTER TABLE person
    ADD COLUMN is_admin boolean NOT NULL DEFAULT false,
    ADD COLUMN is_teacher boolean NOT NULL DEFAULT false,
    ADD COLUMN is_student boolean NOT NULL DEFAULT false,
    ADD COLUMN is_parent boolean NOT NULL DEFAULT false,
    DROP COLUMN role;
