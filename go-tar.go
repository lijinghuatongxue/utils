package main

import (
	"archive/tar"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// 打包多个文件为1个文件
func main() {
	// ----------------- 打包 -------------
	// 目标打包压缩后文件
	dst := "./tmp/a.tar"
	// 目标被打包文件1
	file1 := "./go.mod"
	// 目标被打包文件2
	file2 := "./README.md"
	if err := Tar([]string{file1, file2}, dst); err != nil {
		logrus.Error(err)
	}
	// --------------- 解包 --------------
	// srcFile：目标tar文件 ，dsrDir ：解压目录
	if Decompression("./data/a.tar", "./data") != nil {
		logrus.Error(err)
	}
}

func Tar(src []string, dst string) error {
	// 创建tar文件
	fw, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fw.Close()

	// 通过fw创建一个tar.Writer
	tw := tar.NewWriter(fw)
	// 如果关闭失败会造成tar包不完整
	defer func() {
		if err := tw.Close(); err != nil {
			logrus.Error(err)
		}
	}()

	for _, fileName := range src {
		fi, err := os.Stat(fileName)
		if err != nil {
			logrus.Error(err)
			continue
		}
		hdr, err := tar.FileInfoHeader(fi, "")
		// 将tar的文件信息hdr写入到tw
		err = tw.WriteHeader(hdr)
		if err != nil {
			logrus.Error(err)
			return err
		}

		// 将文件数据写入
		fs, err := os.Open(fileName)
		if err != nil {
			logrus.Error(err)
			return err
		}
		if _, err = io.Copy(tw, fs); err != nil {
			logrus.Error(err)
			return err
		}
		fs.Close()
	}
	return nil
}

func Decompression(srcFile, dsrDir string) error {
	// 打开 tar 包
	fr, err := os.Open(srcFile)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer fr.Close()

	tr := tar.NewReader(fr)
	for hdr, err := tr.Next(); err != io.EOF; hdr, err = tr.Next() {
		if err != nil {
			logrus.Error(err)
			return err
		}
		// 读取文件信息
		fi := hdr.FileInfo()

		// 创建一个空文件，用来写入解包后的数据
		fw, err := os.Create(dsrDir + "/" + fi.Name())
		if err != nil {
			logrus.Error(err)
			return err
		}
		logrus.Infof("Decompression File->", dsrDir+"/"+fi.Name())
		if _, err := io.Copy(fw, tr); err != nil {
			logrus.Error(err)
			return err
		}
		os.Chmod(fi.Name(), fi.Mode().Perm())
		fw.Close()
	}
	return nil
}
