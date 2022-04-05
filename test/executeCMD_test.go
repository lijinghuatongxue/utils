package test

import (
	"fmt"
	"github.com/lijinghuatongxue/utils"
	"testing"
)

func TestExecuteCMD(t *testing.T) {
	cmd := fmt.Sprintf("%s %s %s %s %s", "mkdir ./tmp", "echo 111", "touch ./tmp/12222", "echo 233 > ./tmp/12222", "cat ./tmp/12222")
	err, _, Stdout := utils.LocalCMD(cmd)
	if err != nil {
		return
	}
	fmt.Println(Stdout)
}
