package common

import "github.com/sirupsen/logrus"

var Logger logrus.Logger

func InitLogger() {
	Logger = *logrus.StandardLogger()
}
