package config

var Cfg Config

type Config struct {
	App          AppConfig         `split_words:"true"` // if deployed to k8s
	ServerConfig Http3ServerConfig `split_words:"true" required:"true"`
	QuicConfig   QuicConfig        `split_words:"true" required:"true"`
	ImageConfig  ImageConfig       `split_words:"true" required:"true"`
	JsonConfig   JsonConfig        `split_words:"true" required:"true"`
	MysqlConfig  MySqlConfig       `split_words:"true" required:"true"`
	EventId      int               `split_words:"true" required:"true"`
}

// AppConfig is a struct that contains application's full name and namespace.
type AppConfig struct {
	FullName  string `split_words:"true" required:"true"`
	Namespace string `split_words:"true" required:"true"`
}

type Http3ServerConfig struct {
	Http3ServerAddress string `split_words:"true" required:"true"`
	Http3ServerPort    int    `split_words:"true" required:"true"`
}

type QuicConfig struct {
	HandshakeIdleTimeoutMs int `split_words:"true" required:"true"`
	MaxIdleTimeoutMs       int `split_words:"true" required:"true"`
	KeepAlivePeriod        int `split_words:"true" required:"true"`
}

type ImageConfig struct {
	ImageDirectory string `split_words:"true" required:"true"`
	ImageExtension string `split_words:"true" required:"true"`
}

type JsonConfig struct {
	Directory string `split_words:"true" required:"true"`
}

type MySqlConfig struct {
	Username        string `split_words:"true" required:"true"`
	Password        string `split_words:"true" required:"true"`
	DatabaseAddress string `split_words:"true" required:"true"`
	DatabaseName    string `split_words:"true" required:"true"`
}
