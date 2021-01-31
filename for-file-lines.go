package main

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// 按行读取文件， 传入文件路径
func ForFileLines(FileName string) error {
	fi, err := os.Open(FileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return err
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		logrus.Info(string(a))
	}
	return nil
}
func main() {
	var err error
	err = ForFileLines("./tmp/test.txt")
	if err != nil {
		logrus.Error(err)
	}
}
