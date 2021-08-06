package utils

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"time"
)

func RemoteCmd(IP, Port, User, CMD, idRsaPath string) (string, bool) {
	// ======================================= ssh ===========================
	user := User
	address := IP
	command := CMD
	port := Port
	key, err := ioutil.ReadFile(idRsaPath)
	if err != nil {
		logrus.Errorf("[util - remote-ssh] | ❌ false|unable to read private key: %v", err)
		return "null", false
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	//logrus.Info(key)
	if err != nil {
		logrus.Errorf("[util - remote-ssh] | ❌ false|unable to parse private key: %v", err)
	}
	//hostKeyCallback, err := kh.New("./data/known_hosts")
	if err != nil {
		logrus.Errorf("[util - remote-ssh] | ❌ false|could not create hostkeycallback function: %v ", err)
		return "null", false
	}

	config := &ssh.ClientConfig{
		User:    user,
		Timeout: time.Second * 2, //2s 超时
		Auth: []ssh.AuthMethod{
			// Add in password check here for moar security.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		//	return nil
		//},
	}
	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", address+":"+port, config)
	if err != nil {
		logrus.Errorf("[util - remote-ssh] | ❌ false|unable to connect: %v", err)
		return "null", false
	}
	defer client.Close()
	ss, err := client.NewSession()
	if err != nil {
		logrus.Errorf("[util - remote-ssh] | ❌ false|unable to create SSH session: %v", err)
		return "null", false
	}
	defer ss.Close()
	// Creating the buffer which will hold the remotly executed command's output.
	//var stdoutBuf bytes.Buffer
	//ss.Stdout = &stdoutBuf
	//err = ss.Run(command)
	//执行远程命令
	combo, err := ss.CombinedOutput(command)
	if err != nil {
		logrus.Errorf("[util - remote-ssh] | ❌ false| Err ===》%v", err)
		return "null", false
	}
	return string(combo), true
}
