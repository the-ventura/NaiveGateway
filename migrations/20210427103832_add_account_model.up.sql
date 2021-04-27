CREATE TABLE IF NOT EXISTS accounts (
  id                 serial,
  uuid               varchar unique,
	available          decimal(12,2),
	blocked            decimal(12,2),
	deposited          decimal(12,2),
	withdrawn          decimal(12,2),
	currency           varchar(3),
	card_name          varchar(40),
	card_type          varchar(6),
	card_number        bigint,
	card_expiry_month  integer,
	card_expiry_year   integer,
	card_security_code integer,
	creation_time      timestamptz
);
