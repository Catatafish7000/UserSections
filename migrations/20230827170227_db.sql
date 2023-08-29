-- +goose Up
-- +goose StatementBegin
CREATE TABLE History
(
    user_id    int,
    section    text,
    created_at timestamp,
    deleted_at timestamp
);
CREATE TABLE Sections
(
    user_id    int,
    section    text,
    created_at timestamp,
    ttl        timestamp

);
CREATE table List
(
    section text unique,
    percentage int
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists Sections;
drop table if exists List;
drop table if exists History;
-- +goose StatementEnd
