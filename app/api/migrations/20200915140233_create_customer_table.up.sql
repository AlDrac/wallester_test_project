CREATE TABLE customers (
  id serial NOT NULL,
  first_name varchar(100) NOT NULL,
  last_name varchar(100) NOT NULL,
  birth_date date NOT NULL,
  gender genders NOT NULL,
  email varchar NOT NULL,
  encrypted_password varchar NULL,
  address varchar(200) NULL,
  active bool NOT NULL DEFAULT true,
  registration_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT users_email_key UNIQUE (email),
  CONSTRAINT users_pkey PRIMARY KEY (id)
);