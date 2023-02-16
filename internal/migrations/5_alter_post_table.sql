
-- +migrate Up

ALTER TABLE post DROP FOREIGN KEY post_ibfk_1;
ALTER TABLE post ADD CONSTRAINT author_id_fk FOREIGN KEY(author_id) REFERENCES authors(id);
-- +migrate Down

ALTER TABLE post DROP FOREIGN KEY author_id_fk;
ALTER TABLE post ADD CONSTRAINT post_ibfk_1 FOREIGN KEY (author_id) REFERENCES users(id);