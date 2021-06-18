CREATE TABLE IF NOT EXISTS chats
(
    id           BIGSERIAL  PRIMARY KEY,
    classroom_id BIGINT     NOT NULL ,
    link         VARCHAR    NOT NULL
);