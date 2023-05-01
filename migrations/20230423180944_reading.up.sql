CREATE TABLE IF NOT EXISTS reading(
    id uuid DEFAULT uuid_generate_v4(),
    translation TEXT,
    japanese TEXT,
    title TEXT,
    PRIMARY KEY (id)
)