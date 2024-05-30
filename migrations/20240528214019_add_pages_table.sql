-- +goose Up
-- +goose StatementBegin
CREATE TABLE pages (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    title TEXT NOT NULL,
    link TEXT NOT NULL,
    content TEXT NOT NULL,
    parent_id INTEGER,
    CONSTRAINT `fk_parent_page` FOREIGN KEY (parent_id) REFERENCES pages (id) ON DELETE SET NULL ON UPDATE NO ACTION
)
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE pages;
-- +goose StatementEnd