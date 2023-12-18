CREATE TABLE IF NOT EXISTS admins (
	id INTEGER PRIMARY KEY,
	name  TEXT,
	password  TEXT,
	UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS challenges (
	id INTEGER PRIMARY KEY,
	name      TEXT,
	points    TEXT,
	solution  TEXT,
	UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS players (
	id SERIAL PRIMARY KEY,
	name  TEXT,
	password  TEXT,
	UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS scoreboards (
	id INTEGER PRIMARY KEY,
	name      TEXT,
	teams     TEXT,
	UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS team (
	id INTEGER PRIMARY KEY,
	name      TEXT,
	players   TEXT,
	UNIQUE(name)
);