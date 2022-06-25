CREATE TABLE IF NOT EXISTS kanji_meaning(
    kanji_id uuid REFERENCES kanji (id) ON UPDATE CASCADE ON DELETE CASCADE,
    meaning_id uuid REFERENCES meaning (id) ON UPDATE CASCADE,
    CONSTRAINT PK_kanji_meaning PRIMARY KEY (kanji_id, meaning_id)
);