package test

import (
	"fmt"
	"github.com/lijinghuatongxue/utils"
	"testing"
)

func TestExecuteCMD(t *testing.T) {
	Stdout, err := utils.ExecuteCMD("./", "ls", []string{"-l"})
	if err != nil {
		return
	}
	fmt.Println(Stdout)
}
