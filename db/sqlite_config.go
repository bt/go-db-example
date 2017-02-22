package db

type Sqlite3Config struct {
	path string
}

func (s *Sqlite3Config) Dialect() string {
	return "sqlite3"
}

func (s Sqlite3Config) Path() string {
	return s.path
}
