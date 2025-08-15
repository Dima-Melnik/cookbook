package utils

import "github.com/sirupsen/logrus"

func SendLog(component string, action string, err error) {
	logrus.WithFields(logrus.Fields{
		"Component": component,
		"Action":    action,
	}).Error(err)
}
