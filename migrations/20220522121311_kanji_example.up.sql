CREATE TABLE IF NOT EXISTS kanji_example(
    kanji_id uuid NOT NULL,
    example_word VARCHAR (50),
    CONSTRAINT pk_kanji_example
        PRIMARY KEY (kanji_id, example_word),
    CONSTRAINT fk_kanji_id 
        FOREIGN KEY(kanji_id)
            REFERENCES kanji(id)
            ON DELETE CASCADE
)