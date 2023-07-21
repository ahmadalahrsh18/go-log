package logrus

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type doSomethingHook struct {
	levels []logrus.Level
	skip   int
}

var (
	_ logrus.Hook = (*doSomethingHook)(nil)
)

func (h *doSomethingHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *doSomethingHook) Fire(entry *logrus.Entry) error {
	fmt.Println("do somethings.")
	return nil
}
