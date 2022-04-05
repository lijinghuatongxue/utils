package test

import (
	"fmt"
	"github.com/lijinghuatongxue/utils"
	"testing"
)

func TestExecuteCMD(t *testing.T) {
	err, _, Stdout := utils.LocalCMD("echo 2222")
	if err != nil {
		return
	}
	fmt.Println(Stdout)
}
