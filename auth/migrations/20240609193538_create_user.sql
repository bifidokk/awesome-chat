-- +goose Up
-- +goose StatementBegin
create table "user" (
    id serial primary key,
    name character varying(255) not null,
    email character varying(255) not null,
    role character varying(255) not null,
    password character varying(255) not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user;
-- +goose StatementEnd
