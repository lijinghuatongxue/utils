package utils

import (
	"fmt"
	"github.com/pkg/sftp"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

func connect(isPasswd bool, user, password, PublicKeysPath, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	pemBytes, err := ioutil.ReadFile(PublicKeysPath)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(pemBytes)
	if err != nil {
		logrus.Error("parse key failed:%v", err)
		return nil, err
	}
	if isPasswd {
		auth = append(auth, ssh.Password(password))
	} else {
		auth = append(auth, ssh.PublicKeys(signer))
	}

	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}

	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}

func SftpUpload(isPasswd bool, username, passwd, sftpServer, PublicKeysPath, LocalFilePath, RemoteDirPath string, sftpPort int) error {
	var (
		err        error
		sftpClient *sftp.Client
	)
	if isPasswd == false {
		passwd = "xx"
	}
	// 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口
	sftpClient, err = connect(isPasswd, username, passwd, PublicKeysPath, sftpServer, sftpPort)
	if err != nil {
		log.Fatal(err)
	}
	defer func(sftpClient *sftp.Client) {
		err := sftpClient.Close()
		if err != nil {
			logrus.Error(err)
			return
		}
	}(sftpClient)
	srcFile, err := os.Open(LocalFilePath)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {
			logrus.Error(err)
			return
		}
	}(srcFile)

	var remoteFileName = path.Base(LocalFilePath)
	dstFile, err := sftpClient.Create(path.Join(RemoteDirPath, remoteFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer func(dstFile *sftp.File) {
		err := dstFile.Close()
		if err != nil {
			logrus.Error(err)
			return
		}
	}(dstFile)

	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}

	logrus.Infof("上传成功 ✅ ｜远程文件位置 -> %s", RemoteDirPath+remoteFileName)
	return nil
}

func SftpDownload(isPasswd bool, username, passwd, sftpServer, PublicKeysPath, LocalDirPath, RemoteFilePath string, sftpPort int) error {
	var (
		err        error
		sftpClient *sftp.Client
	)
	if isPasswd == false {
		passwd = "xx"
	}
	// 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口
	sftpClient, err = connect(isPasswd, username, passwd, PublicKeysPath, sftpServer, sftpPort)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer func(sftpClient *sftp.Client) {
		err := sftpClient.Close()
		if err != nil {
			logrus.Error(err)
			return
		}
	}(sftpClient)

	srcFile, err := sftpClient.Open(RemoteFilePath)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer func(srcFile *sftp.File) {
		err := srcFile.Close()
		if err != nil {
			logrus.Error(err)
			return
		}
	}(srcFile)

	var localFileName = path.Base(RemoteFilePath)
	dstFile, err := os.Create(path.Join(LocalDirPath, localFileName))
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer func(dstFile *os.File) {
		err := dstFile.Close()
		if err != nil {
			logrus.Error(err)
			return
		}
	}(dstFile)

	if _, err = srcFile.WriteTo(dstFile); err != nil {
		log.Fatal(err)
		return err
	}
	logrus.Infof("下载成功 ✅ ｜文件位置 -> %s", LocalDirPath+path.Base(RemoteFilePath))
	return nil
}
