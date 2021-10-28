package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	_ "html/template"
)

var mailConfig = map[string]string{
	"MAIL_SMTP_HOST": "smtp.163.com",           // 服务器地址
	"MAIL_USER":      "tashilijinghua@163.com", // 发件人的邮箱地址
	"MAIL_PASS":      "KHIKROHZIQSYUABY",       // 口令
	"ME":             "",                       // 不填
}

const MailMess = // 定义邮件表格样式
`<style>
		table {
			border-collapse: collapse;
		}
		th {
			background-color: #007fff;
			color: white;
		}
		table, th, td {
			border: 1px solid black;
			padding: 5px;
			text-align: left;
		}

		</style>
		<table>
			<tr>
				<th>集群号</th>
				<th>所属业务</th>
				<th>源库</th>
				<th>源表</th>
				<th>目标库</th>
				<th>目标表</th>
				<th>插入行数</th>
				<th>删除行数</th>
				<th>avg_row_length</th>
				<th>消耗时间</th>
			</tr>`

func FormShow() { // 调用接口传过来的数据，拼接表格
	var mailMessage = MailMess
	mailMessage = MailMess + fmt.Sprintf("<td>%v</td><td>%v</td><td>%v</td><td>%v</td><td>%v</td><td>%v</td><td>%v</td><td>%v</td><td>%v</td><td>%v</td></tr>",
		"111", "222", "333", "444", "555", "666", "777", "888", "999", "1010")
	mailMessage = mailMessage + "</table><br></br><br>Send By tashilijinghua@163.com</br>（自动发送请勿回复）"
	DayEmail(mailMessage) //拼接完成后，调用函数发送邮件
}

func DayEmail(mailMess string) { // 发送邮件
	m := gomail.NewMessage()
	m.SetHeader("From", mailConfig["MAIL_USER"]) //发件人
	m.SetHeader("To", "598820383@qq.com")        //收件人
	//m.SetAddressHeader("Cc", "test@126.com", "test")     //抄送人
	m.SetHeader("Subject", "golang 发送邮件测试") //邮件标题
	m.SetBody("text/html", mailMess)        //邮件内容
	m.Attach("./tcpScanPort.go")            //邮件附件
	d := gomail.Dialer{Host: mailConfig["MAIL_SMTP_HOST"], Port: 25, Username: mailConfig["MAIL_USER"], Password: mailConfig["MAIL_PASS"], SSL: false}
	//邮件发送服务器信息,使用授权码而非密码
	if err := d.DialAndSend(m); err != nil {
		logrus.Error(err)
	}
}
func main() {
	FormShow() // 传进去的是[]struct的切片
}
