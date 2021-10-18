package logrgorm2_test

import (
	stdlog "log"
	"os"
	"testing"

	"github.com/go-logr/stdr"
	"github.com/vchitai/logrgorm2"

	"gorm.io/gorm"
)

func TestExample(t *testing.T) {
	stdr.SetVerbosity(1)
	log := stdr.NewWithOptions(stdlog.New(os.Stderr, "", stdlog.LstdFlags), stdr.Options{LogCaller: stdr.All})
	log = log.WithName("logr")
	logger := logrgorm2.New(log)
	db, _ := gorm.Open(nil, &gorm.Config{Logger: logger})

	// do stuff normally
	var _ = db // avoid "unused variable" warn
}
