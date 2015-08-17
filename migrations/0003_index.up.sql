CREATE UNIQUE INDEX term_lower_name_year_unique ON term (lower(name),school_year);
CREATE UNIQUE INDEX account_lower_email_unique ON account (lower(email));
CREATE UNIQUE INDEX level_lower_name_unique ON level (lower(name));
