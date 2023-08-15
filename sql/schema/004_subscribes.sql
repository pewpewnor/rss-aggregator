-- +goose Up

CREATE TABLE subscribes (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE(user_id, feed_id)
);

-- +goose Down

DROP TABLE subscribes;