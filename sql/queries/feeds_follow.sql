-- name: CreateFeedFollow :one
WITH insert_feed_follow AS (
    INSERT INTO feed_follows ( created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4
    )
    RETURNING *
)

SELECT  
    i.id,
    i.created_at,
    i.updated_at,
    u.name AS user_name,
    f.name AS feed_name
FROM insert_feed_follow i 
LEFT JOIN feeds f ON f.id = i.feed_id
LEFT JOIN users u ON u.id = i.user_id; 

-- name: GetFeedsForUser :many
SELECT  
    f.name AS feed_name
FROM feed_follows ff 
LEFT JOIN feeds f ON f.id = ff.feed_id
WHERE ff.user_id = $1;
