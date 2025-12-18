package setting

type Config struct {
	Mysql  MySQLSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"log"`
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

type LoggerSetting struct {
	LogLevel   string `mapstructure:"log_level"`
	FileName   string `mapstructure:"file_log_name"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}
