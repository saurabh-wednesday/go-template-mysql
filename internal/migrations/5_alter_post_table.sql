
-- +migrate Up

ALTER TABLE post DROP FOREIGN KEY post_ibfk_1;
ALTER TABLE post ADD FOREIGN KEY(author_id) REFERENCES authors(id);
-- +migrate Down
DROP TABLE post;  
