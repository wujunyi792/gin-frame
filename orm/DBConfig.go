package orm

type DBConfig struct {
	Type       string       `toml:"type"`
	Address    string       `toml:"address"`
	AuthConfig DBAuthConfig `toml:"auth"`
	DBName     string       `toml:"dbName"`
}

type DBAuthConfig struct {
	UserName string `toml:"userName"`
	Password string `toml:"password"`
}
