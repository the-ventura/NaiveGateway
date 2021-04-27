CREATE TABLE IF NOT EXISTS transactions(
  id            serial,
	from_id       varchar,
	to_id         varchar,
	amount        decimal(12,2),
	status        varchar(10),
	description   varchar,
	currency      varchar(3),
  uuid          varchar primary key,
	creation_time timestamptz
);
