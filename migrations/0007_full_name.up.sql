CREATE FUNCTION display_name(rec person)
  RETURNS text
  STABLE
  LANGUAGE SQL
  COST 5
AS $$
  SELECT
    $1.first_name || ' ' ||
    CASE
      WHEN $1.middle_name IS NULL THEN ''
      ELSE $1.middle_name || ' '
    END || $1.last_name;
$$;
