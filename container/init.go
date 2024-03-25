package container

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

func readUserCommand() []string {
	pipe := os.NewFile(uintptr(3), "pipe")
	bs, err := io.ReadAll(pipe)
	if err != nil {
		logrus.Errorf("read pipe,err: %v", err)
		return nil
	}

	msg := string(bs)
	return strings.Split(msg, " ")
}
