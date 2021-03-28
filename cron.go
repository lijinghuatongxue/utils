package main

import (
	"github.com/sirupsen/logrus"
)

func FuncExample(Str string) {
	logrus.Info(Str)
}
func main() {
	// todo
	FuncExample("233")
}
