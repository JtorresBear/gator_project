-- name: Create_Post :one
INSERT INTO posts(id,created_at,updated_at,title,url,description,published_at,feed_id)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;

-- name: Get_Posts_for_User :many
SELECT *
FROM posts
ORDER BY published_at DESC
LIMIT $1;

-- name: Reset_Posts :exec
DELETE FROM posts;