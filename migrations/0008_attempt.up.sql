ALTER TABLE attempt
    ADD COLUMN score_number numeric,
    DROP COLUMN average;


--
-- SELECT
-- enrollment.*,
-- person.display_name AS "person.display_name",
-- assignment.assignment_id AS "assignment.assignment_id",
-- assignment.name AS "assignment.name",
-- assignment.max_score AS "assignment.max_score",
-- assignment.weight AS "assignment.weight",
-- assignment.group_name AS "assignment.group_name",
-- attempt.score,
-- ((assignment.max_score/attempt.score::INT)*10)
-- FROM enrollment
-- INNER JOIN person USING(person_id)
-- INNER JOIN (SELECT assignment.*, assignment_group.name AS "group_name", assignment_group.weight AS "weight" FROM assignment JOIN assignment_group USING(group_id)) AS assignment USING(course_id,term_id)
-- LEFT OUTER JOIN (SELECT * FROM attempt WHERE attempt_id IN (SELECT max(attempt_id) FROM attempt GROUP BY person_id)) AS attempt USING(assignment_id, person_id)
