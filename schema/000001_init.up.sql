CREATE TABLE IF NOT EXISTS users
(
    id         serial primary key,
    username   varchar(255) not null,
    first_name  varchar(255) not null,
    last_name  varchar(255) not null,
    role       varchar(255) not null,
    password   varchar(255) not null,
    created_at timestamp not null default current_timestamp,
    deleted_at timestamp,
    updated_at timestamp not null default current_timestamp,
    constraint users_username_idx unique
        (username)
  );