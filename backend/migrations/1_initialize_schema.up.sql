BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table users
(
    id         uuid          default uuid_generate_v4(),
    first_name text not null,
    last_name  text not null default '',
    email      text not null unique,
    avatar_url text,
    address    text,
    is_admin   bool not null,
    primary key (id)
);

COMMIT;
