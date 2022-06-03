CREATE TABLE IF NOT EXISTS kanji_example(
    id uuid DEFAULT uuid_generate_v4(),
    kanji_id uuid NOT NULL,
    example_word VARCHAR (50),
    PRIMARY KEY (id),
    CONSTRAINT fk_kanji_id 
        FOREIGN KEY(kanji_id)
            REFERENCES kanji(id)
            ON DELETE CASCADE
)