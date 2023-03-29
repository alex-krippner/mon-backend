CREATE TABLE IF NOT EXISTS vocabulary(
  id uuid DEFAULT uuid_generate_v4(),
  vocab VARCHAR (50) NOT NULL,
  kanji VARCHAR (50),
  vocab_rating INT,
  username VARCHAR (50),
  definitions TEXT,
  example_sentences TEXT,
  parts_of_speech TEXT,
  PRIMARY KEY (id)
);