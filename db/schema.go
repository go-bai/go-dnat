package db

var schema = `
CREATE TABLE IF NOT EXISTS rule (
	id INTEGER PRIMARY KEY,
	iface TEXT,
	port INTEGER,
	dest TEXT,
	created_at DATETIME,
	UNIQUE(iface, port)
)
`
