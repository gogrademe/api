CREATE TABLE announcement (
  id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  title text,
  body text,
  posted_date timestamp,
  person_id integer NOT NULL
);

CREATE TABLE assignment (
  id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  name text NOT NULL,
  max_score integer,
  due_date timestamp,
  group_id integer NOT NULL,
  term_id integer NOT NULL,
  course_id integer NOT NULL
);


CREATE TABLE assignment_attempt (
  id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  score text,
  grade_average decimal,
  person_id integer NOT NULL,
  assignment_id integer NOT NULL
);

CREATE TABLE assignment_group (
  id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  name text NOT NULL,
  weight decimal,
  course_id integer NOT NULL,
  term_id integer NOT NULL
);

CREATE TABLE course (
  id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  name text NOT NULL,
  level_id integer not null,
  max_students integer
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
  id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  course_id integer NOT NULL UNIQUE,
  person_id integer NOT NULL UNIQUE,
  term_id integer NOT NULL UNIQUE
);

CREATE TABLE person (
  id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  first_name text NOT NULL,
  middle_name text,
  last_name text NOT NULL,
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
  id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  token text NOT NULL,
  expires_at timestamp NOT NULL,
  account_id integer NOT NULL
);

CREATE TABLE term (
  id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  name text NOT NULL,
  school_year int NOT NULL
);

CREATE TABLE level (
  id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  name text NOT NULL
);

CREATE TABLE account (
  id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  archived_at timestamp NULL,
  email text NOT NULL UNIQUE,
  role text,
  hashed_password text NOT NULL,
  activation_token text,
  disabled boolean,
  person_id integer NOT NULL
);
