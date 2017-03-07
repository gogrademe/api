CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_attempt_modtime BEFORE INSERT OR UPDATE ON attempt FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
CREATE TRIGGER update_person_modtime BEFORE INSERT OR UPDATE ON person FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
CREATE TRIGGER update_assignment_modtime BEFORE INSERT OR UPDATE ON assignment FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
CREATE TRIGGER update_course_modtime BEFORE INSERT OR UPDATE ON course FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
