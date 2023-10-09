CREATE TABLE cas_records (
	id serial4 primary key,
	name VARCHAR(255),
  	marks jsonb,
  	created TIMESTAMPTZ
);