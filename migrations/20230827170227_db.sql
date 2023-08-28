-- +goose Up
-- +goose StatementBegin
CREATE TABLE Sections
(
    user_id int,
    section text,
    created_at timestamp,
    ttl timestamp

);
CREATE table List
(
        section text unique
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists Sections;
drop table if exists List
-- +goose StatementEnd
