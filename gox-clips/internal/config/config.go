package config

type AppConfig struct {
    Port      int
    DB_URL    string
    DebugMode bool
}

func Config() *AppConfig {
	return &AppConfig{
		Port:      8080,
		DB_URL:    "sqlite3://app.db",
		DebugMode: true,
	}
}
