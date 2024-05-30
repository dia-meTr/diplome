BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table users
(
    id         uuid          default uuid_generate_v4(),
    first_name text not null,
    last_name  text not null default '',
    email      text not null unique,
    user_role  text not null,
    primary key (id)
);

COMMIT;
