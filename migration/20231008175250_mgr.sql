-- +goose Up
-- +goose StatementBegin
create table articles
(
    id           serial
        constraint articles_pk
            primary key,
    name_article text    not null,
    rating       integer not null,
    created_at   timestamp
);

create table post
(
    id        serial
        constraint post_pk
            primary key,
    id_author integer not null
        constraint post
            references articles
            on delete cascade,
    name_post text    not null,
    sales     integer not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table post;
drop table articles;
-- +goose StatementEnd