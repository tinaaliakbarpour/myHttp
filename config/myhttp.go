package config

import "time"

type MyHttp struct {
	TimeOut time.Duration `yaml:"myHttp.timeout"`
}
