package meUtils

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
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
		logrus.Error("[util - remote-ssh] | ❌ false|unable to read private key: %v", err)
		return "null", false
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	//logrus.Info(key)
	if err != nil {
		logrus.Error("[util - remote-ssh] | ❌ false|unable to parse private key: %v", err)
	}
	//hostKeyCallback, err := kh.New("./data/known_hosts")
	if err != nil {
		logrus.Error("[util - remote-ssh] | ❌ false|could not create hostkeycallback function: ", err)
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
		logrus.Error("[util - remote-ssh] | ❌ false|unable to connect: %v", err)
		return "null", false
	}
	defer client.Close()
	ss, err := client.NewSession()
	if err != nil {
		logrus.Error("[util - remote-ssh] | ❌ false|unable to create SSH session: ", err)
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
		log.Fatal("远程执行cmd 失败", err)
	}
	if err != nil {
		logrus.Error("[util - remote-ssh] | ❌ false| Err ===》%s", err)
		return "null", false
	}
	return string(combo), true
}

// 远程执行ssh指令
// Remote execution of SSH instruction
//func main() {
//	cmd := "docker ps"
//	ip := "192.168.0.123"
//	port := "22"
//	user := "root"
//	RemoteCmdOutput, err := RemoteCmd(ip, port, user, cmd)
//	if err != true {
//		logrus.Error("[util - remote-ssh] | ❌ false| Err ===》%s", err)
//		return
//	}
//	logrus.Infof("cmd -> %s |Output -> \n%s", cmd, RemoteCmdOutput)
//}
