CREATE TABLE IF NOT EXISTS chats
(
    id           BIGSERIAL  UNIQUE PRIMARY KEY,
    classroom_id BIGINT     NOT NULL ,
    link         VARCHAR    NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS classroom_and_link_unique
    ON chats (classroom_id, link);