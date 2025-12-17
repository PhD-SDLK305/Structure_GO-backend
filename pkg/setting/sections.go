package setting

type Config struct {
	Mysql MySQLSetting `mapstructure:"mysql"`
}

type MySQLSetting struct {
	Host            string `mapstructure="host"`
	Port            int    `mapstructure="port"`
	Username        string `mapstructure="username"`
	Passwrod        string `mapstructure="passwrod"`
	Dbname          string `mapstructure="dbname"`
	MaxIdleConns    int    `mapstructure="maxIdleConns"`
	MaxOpenConns    int    `mapstructure="maxOpenConns"`
	ConnMaxLifetime int    `mapstructure="connMaxLifetime"`
}
