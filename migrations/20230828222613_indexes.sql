-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX unique_pair_udx ON Sections(user_id,section);
create unique index unique_history_idx on history(user_id,section);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
