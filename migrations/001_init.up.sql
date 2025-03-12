CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS songs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "group" VARCHAR(255) NOT NULL,
    song VARCHAR(255) NOT NULL,
    release_date DATE NOT NULL,
    link TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS lyrics (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    song_id UUID REFERENCES songs(id) ON DELETE CASCADE,
    verse_number INTEGER NOT NULL,
    text TEXT NOT NULL
);