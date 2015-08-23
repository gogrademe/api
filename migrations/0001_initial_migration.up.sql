CREATE TABLE announcement (
  announcement_id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  title text,
  body text,
  posted_date timestamp,
  person_id integer NOT NULL
);

CREATE TABLE assignment (
  assignment_id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  name text NOT NULL,
  max_score integer,
  due_date timestamp,
  assignment_group_id integer NOT NULL,
  term_id integer NOT NULL,
  course_id integer NOT NULL
);


CREATE TABLE attempt (
  attempt_id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  score text,
  grade_average decimal,
  person_id integer NOT NULL,
  assignment_id integer NOT NULL
);

CREATE TABLE assignment_group (
  assignment_group_id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  name text NOT NULL,
  weight decimal,
  course_id integer NOT NULL,
  term_id integer NOT NULL
);

CREATE TABLE course (
  course_id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  name text NOT NULL,
  level_id integer not null,
  max_students integer
);

CREATE TABLE course_term (
  course_id integer,
  term_id integer,
  PRIMARY KEY (course_id, term_id)
);

CREATE TABLE email_confirmation (
  id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  used_on timestamp NULL,
  account_id integer NOT NULL
);

CREATE TABLE enrollment (
  enrollment_id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  course_id integer NOT NULL,
  person_id integer NOT NULL,
  term_id integer NOT NULL,
  UNIQUE (course_id, person_id, term_id)
);

CREATE TABLE person (
  person_id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  first_name text NOT NULL,
  middle_name text,
  last_name text NOT NULL,
  role integer NOT NULL,
  grade_level text
);

CREATE TABLE contact_info (
  id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  phone_number text,
  contact_email text
);

CREATE TABLE session (
  session_id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  token text NOT NULL,
  expires_at timestamp NOT NULL,
  account_id integer NOT NULL
);

CREATE TABLE term (
  term_id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  name text NOT NULL,
  school_year int NOT NULL
);

CREATE TABLE level (
  level_id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  name text NOT NULL
);

CREATE TABLE account (
  account_id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  email text NOT NULL,
  hashed_password text NOT NULL,
  activation_token text,
  disabled boolean,
  person_id integer NOT NULL UNIQUE
);
