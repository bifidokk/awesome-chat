-- +goose Up
-- +goose StatementBegin
create table chat (
    id serial primary key,
    usernames jsonb,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chat;
-- +goose StatementEnd
