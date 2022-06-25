CREATE TABLE IF NOT EXISTS meaning(
    id uuid DEFAULT uuid_generate_v4(),
    meaning VARCHAR(50),
    PRIMARY KEY (id)
);