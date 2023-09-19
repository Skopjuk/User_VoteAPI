CREATE TABLE IF NOT EXISTS votes
(
    id                      serial primary key,
    user_id                 bigint,
    rated_user_id           bigint,
    vote                    int,
    created_at              timestamp not null default current_timestamp,
    updated_at              timestamp not null default current_timestamp,
    constraint fk_user_id
        foreign key(user_id)
            references users(id),
    constraint fk_rated_user_id
        foreign key(rated_user_id)
            references users(id)
);

CREATE TABLE IF NOT EXISTS ratings
(
    id serial primary key,
    user_id bigint,
    rating int,
    created_at              timestamp not null default current_timestamp,
    updated_at              timestamp not null default current_timestamp,
    constraint fk_user_id
        foreign key(user_id)
            references users(id)
)