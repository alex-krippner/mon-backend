CREATE TABLE IF NOT EXISTS reading(
    id uuid DEFAULT uuid_generate_v4(),
    japanese TEXT,
    english_translation TEXT,
    PRIMARY KEY (id)
)