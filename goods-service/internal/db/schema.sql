CREATE TABLE IF NOT EXISTS ads (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    title VARCHAR(100) NOT NULL CHECK (char_length(title) > 0),
    description TEXT NOT NULL CHECK (char_length(description) <= 1000),
    image_url VARCHAR(500) NOT NULL CHECK (image_url ~* '^https?://'),
    price NUMERIC(10, 2) NOT NULL CHECK (price >= 0),

    author_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_ads_created_at ON ads (created_at DESC);
CREATE INDEX IF NOT EXISTS idx_ads_price ON ads (price);
