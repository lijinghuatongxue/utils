package utils

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func RMUtils(path string, onlyEmpty bool) error {
	if IsDir(path) {
		if onlyEmpty {
			fileList, err := FindDirAllFileNoSuffix(path)
			if err != nil {
				logrus.Error(err)
			}
			for _, f := range fileList {
				err := os.Truncate(f, 0)
				if err != nil {
					log.Fatal(err)
					return err
				}
			}
		} else {
			err := os.RemoveAll(path)
			if err != nil {
				logrus.Error(err)
				return err
			}
		}
	} else {
		if onlyEmpty {
			err := os.Truncate(path, 0)
			if err != nil {
				log.Fatal(err)
				return err
			}
		} else {
			err := os.RemoveAll(path)
			if err != nil {
				logrus.Error(err)
				return err
			}
		}

	}
	return nil
}
func RM(path string, onlyFileEmpty bool) error {
	err := RMUtils(path, onlyFileEmpty)
	if err != nil {
		return err
	}
	return nil
}
