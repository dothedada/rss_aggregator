-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: ListFeeds :many
SELECT 
	feeds.name as "feed name", 
	url, 
	users.name as "suscribed user"
FROM feeds
JOIN users ON feeds.user_id = users.id;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow as (
	INSERT INTO feed_follow (id, created_at, updated_at, user_id, feed_id) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING *
) 
SELECT
	inserted_feed_follow.*,
	feeds.name as feed_name,
	users.name as user_name
FROM inserted_feed_follow
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id
INNER JOIN users ON users.id = feeds.user_id;

-- name: GetFeedByUrl :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeedsFollowedByUser :many
SELECT feeds.name as feed_name, users.name as user_name 
FROM feeds
INNER JOIN feed_follow ON feeds.id = feed_follow.feed_id
INNER JOIN users ON feed_follow.user_id = users.id
WHERE user.name = $1;
