CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY,
    username VARCHAR(60)
);

CREATE TABLE IF NOT EXISTS subscribe (
    id INTEGER PRIMARY KEY,
    user_id INTEGER,
    source VARCHAR(100),
    tag VARCHAR(100)
);
