CREATE TABLE IF NOT EXISTS exoplanets (
        id SERIAL PRIMARY KEY,
        "name" TEXT,
        "description" TEXT,
        distance DOUBLE PRECISION,
        radius DOUBLE PRECISION,
        mass DOUBLE PRECISION,
        "type" TEXT
)