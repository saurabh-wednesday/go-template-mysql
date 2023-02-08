
-- +migrate Up
CREATE TABLE post (
			id INT AUTO_INCREMENT PRIMARY KEY,
			author_id int,
            FOREIGN KEY (author_id) REFERENCES users(id),
			title VARCHAR(200)  NOT NULL,
            body VARCHAR(4000) NOT NULL,
			created_at TIMESTAMP,
			updated_at TIMESTAMP,
			deleted_at TIMESTAMP
		);
-- +migrate Down
DROP TABLE post;