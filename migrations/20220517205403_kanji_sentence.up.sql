CREATE TABLE IF NOT EXISTS kanji_sentence(
    kanji_id uuid NOT NULL,
    example_sentence TEXT,
    CONSTRAINT pk_kanji_sentence 
        PRIMARY KEY (kanji_id, example_sentence),
    CONSTRAINT fk_kanji_id
        FOREIGN KEY(kanji_id)
            REFERENCES kanji(id)
            ON DELETE CASCADE
);