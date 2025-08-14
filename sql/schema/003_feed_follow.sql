-- +goose Up
CREATE TABLE feed_follow (
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	user_id UUID NOT NULL,
	feed_id UUID NOT NULL,
	CONSTRAINT unique_user_feed 
	UNIQUE(user_id, feed_id),
	CONSTRAINT fk_user_id
	FOREIGN KEY (user_id)
	REFERENCES users(id)
	ON DELETE CASCADE,
	CONSTRAINT fk_feed_id
	FOREIGN KEY (feed_id)
	REFERENCES feeds(id)
	ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feed_follow;
