package logger

import (
	log "github.com/sirupsen/logrus"

	"github.com/bernardolm/step-task/pkg/infrastructure/errors"
)

func Entry(err error) *log.Entry {
	en := log.NewEntry(log.StandardLogger())

	if err != nil {
		en = en.WithField("cause", err)

		if appErr, ok := err.(*errors.AppError); ok {
			en = en.
				WithField("cause", appErr.Message()).
				WithFields(appErr.Fields())
		}
	}

	return en
}
