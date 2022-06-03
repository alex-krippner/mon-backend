CREATE TABLE IF NOT EXISTS vocabulary_definition(
    id uuid DEFAULT uuid_generate_v4(), 
    vocab_id uuid NOT NULL,
    def TEXT,
    PRIMARY KEY (id),
    CONSTRAINT fk_vocab_id
        FOREIGN KEY(vocab_id)
            REFERENCES vocabulary(id)
            ON DELETE CASCADE
);