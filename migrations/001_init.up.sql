CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS groups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS songs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    group_id UUID NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    song VARCHAR(255) NOT NULL,
    release_date DATE NOT NULL,
    link TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS lyrics (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    song_id UUID NOT NULL REFERENCES songs(id) ON DELETE CASCADE,
    verse_number INTEGER NOT NULL,
    text TEXT NOT NULL
);

CREATE INDEX idx_songs_group_id ON songs(group_id);
CREATE INDEX idx_songs_release_date ON songs(release_date);
CREATE INDEX idx_lyrics_song_id ON lyrics(song_id);
