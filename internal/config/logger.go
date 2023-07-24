package config

import (
	"go.uber.org/zap"
)

type apiLogger struct {
	log *zap.Logger
}

func NewApiLogger() *apiLogger {
	//var l = zap.New()
	//l. = os.Stdout
	//l.SetFormatter(&logrus.TextFormatter{
	//	FullTimestamp: true,
	//})

	l := zap.Must(zap.NewProduction())

	defer l.Sync()

	return &apiLogger{
		log: l,
	}
}
