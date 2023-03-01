package global

type ServerConf struct {
	Port   int    `mapstructure:"port"`
	RunEnv string `mapstructure:"runEnv"`
}

type MysqlConf struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DB              string `mapstructure:"database"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
	InitSQL         string `mapstructure:"initSQL"`
}

type RedisConf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"database"`
}

type JwtConf struct {
	SigningKey string `mapstructure:"signingKey"`
	ExpireTime int    `mapstructure:"expireTime"`
}

type UploadConf struct {
	Path string `mapstructure:"path"`
}

type MailConf struct {
	SmtpServer string `mapstructure:"smtp"`
	AuthCode   string `mapstructure:"authCode"`
	Sender     string `mapstructure:"sender"`
}

type Config struct {
	ServerConf *ServerConf `mapstructure:"server"`
	MysqlConf  *MysqlConf  `mapstructure:"mysql"`
	RedisConf  *RedisConf  `mapstructure:"redis"`
	JwtConf    *JwtConf    `mapstructure:"jwt"`
	UploadConf *UploadConf `mapstructure:"upload"`
	MailConf   *MailConf   `mapstructure:"mail"`
}
