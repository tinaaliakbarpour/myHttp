package config

var (
	// Global config
	Confs = Config{}
)

type config interface {
	Set(key string, query []byte) error
	Get() Config
	SetDebug(bool)
	GetDebug() bool
	Load(path string) error
}

// Config is base of configs we need for project
type Config struct {
	Debug    bool     // if true we run on debug mode
	MyHttp  MyHttp  `yaml:"myHttp"`
}
