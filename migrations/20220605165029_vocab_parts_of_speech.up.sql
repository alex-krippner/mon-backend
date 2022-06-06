CREATE TABLE IF NOT EXISTS vocab_part_of_speech(
    vocab_id uuid REFERENCES vocabulary (id) ON UPDATE CASCADE ON DELETE CASCADE,
    part_of_speech_id uuid REFERENCES parts_of_speech (id) ON UPDATE CASCADE,
    CONSTRAINT PK_vocab_part_of_speech PRIMARY KEY (vocab_id, part_of_speech_id)
);