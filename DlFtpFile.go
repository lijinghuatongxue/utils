package meUtils

import (
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"os"
	"time"
)

func DlFtpFile(FtpPort, FtpServerAddr, SavePath, DlFilePath string) {
	c, err := ftp.Dial(FtpServerAddr+":"+FtpPort, ftp.DialWithTimeout(10*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("demo", "password")
	if err != nil {
		log.Fatal(err)
	}

	//c.ChangeDir("./data")

	res, err := c.Retr(DlFilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	outFile, err := os.Create(SavePath)
	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, res)
	if err != nil {
		log.Fatal(err)
	}
}

//func main()  {
//	DlFtpFile("21","test.rebex.net","./data/test-ftp.png","/pub/example/mail-editor.png")
//}
