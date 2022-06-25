CREATE TABLE IF NOT EXISTS parts_of_speech(
    id uuid DEFAULT uuid_generate_v4(),
    part_of_speech VARCHAR(50),
    PRIMARY KEY (id)
);