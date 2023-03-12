create extension if not exists "citext" with schema "public";

create table users (
  id uuid not null primary key,
  email citext not null,
  password_hash text not null
);

create unique index users_email_idx on users (email);
