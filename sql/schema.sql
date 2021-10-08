CREATE TABLE links
(
    id            BIGSERIAL PRIMARY KEY,
    short_link    text NOT NULL,
    original_link text NOT NULL
);