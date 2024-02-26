package envconf

import (
	"github.com/kelseyhightower/envconfig"
)

type Spec struct {
	AppName   string `envconfig:"APP_NAME" default:"try-gorm"`
	Port      string `envconfig:"PORT" default:"8080"`
	SqliteDsn string `envconfig:"SQLITE_DSN" default:"test.db"`
}

func New() *Spec {
	var s Spec
	envconfig.MustProcess("", &s)
	return &s
}
