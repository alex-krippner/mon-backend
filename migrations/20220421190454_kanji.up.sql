CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS kanji(
  id uuid DEFAULT uuid_generate_v4(),
  kanji VARCHAR (50) NOT NULL,
  on_reading VARCHAR (50),
  kun_reading VARCHAR (50),
  kanji_rating INT,
  username VARCHAR (50),
  meanings TEXT,
  example_sentences TEXT,
  example_words TEXT,
  PRIMARY KEY (id)
);