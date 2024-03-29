CREATE TABLE IF NOT EXISTS admins (
	id SERIAL PRIMARY KEY,
	name  	  TEXT,
	password  TEXT,
	UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS campaigns (
    id SERIAL PRIMARY KEY,
	name      TEXT,
	UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS teams (
	id INTEGER PRIMARY KEY,
	name      TEXT,
	UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS players (
	id SERIAL PRIMARY KEY,
	name  TEXT,
	password  TEXT,
	team INT REFERENCES teams(id),
	UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS challenges (
    id SERIAL PRIMARY KEY,
	name      TEXT,
	points    TEXT,
	campaign  INT REFERENCES campaigns(id),
	team      INT REFERENCES teams(id),
	solution  TEXT,
	UNIQUE(name)
);
