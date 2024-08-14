CREATE TABLE batting (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    team VARCHAR(100) NOT NULL,
    year INT NOT NULL CHECK (year >= 1900 AND year <= 2100),
    at_bat INT NOT NULL CHECK (at_bat >= 0),
    hit INT NOT NULL CHECK (hit >= 0),
    double INT NOT NULL CHECK (double >= 0),
    triple INT NOT NULL CHECK (triple >= 0),
    home_run INT NOT NULL CHECK (home_run >= 0),
    ball_on_base INT NOT NULL CHECK (ball_on_base >= 0),
    hit_by_pitch INT NOT NULL CHECK (hit_by_pitch >= 0),
    CONSTRAINT unique_constraint UNIQUE (name, team, year)
);

CREATE TABLE leaderboard (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    team VARCHAR(100) NOT NULL,
    score FLOAT NOT NULL
);