-- +migrate Up
CREATE TABLE authors (
				id INT AUTO_INCREMENT PRIMARY KEY,
				first_name VARCHAR(400)  NOT NULL,
				last_name VARCHAR(400)  NOT NULL,
				created_at TIMESTAMP,
				updated_at TIMESTAMP,
				deleted_at TIMESTAMP
			);

-- +migrate Down
DROP TABLE authors;