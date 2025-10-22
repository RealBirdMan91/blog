-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS posts (
  id         uuid PRIMARY KEY NOT NULL,
  body       TEXT NOT NULL,
  author_id  uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  likes      integer NOT NULL DEFAULT 0 CHECK (likes >= 0),
  views      integer NOT NULL DEFAULT 0 CHECK (views >= 0),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_posts_author_id ON posts(author_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_posts_author_id;
DROP TABLE IF EXISTS posts;
-- +goose StatementEnd
