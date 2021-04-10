package meUtils

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"os"
	"text/template"
)

// tpl demo
//[program:{{.ProjectName}}]
//directory = {{.ProjectName}}
//command = {{.CMD}}
//user=root
//numprocs=1
//stopsignal=KILL
//startretries=5
//autostart=true
//redirect_stderr=true
//stdout_logfile = /opt/supervisord/var/log/{{.ProjectName}}.log

func GenerateTpl(WFilePath, TplPath string, TplStruct interface{}, DetailedOutput bool) (error, string) {
	// 解析模版
	tmpl, err := template.ParseFiles(TplPath)
	if err != nil {
		logrus.Errorf("[util - 模版解析] ｜ ❌ false| err ===》%s ", err)
		return err, WFilePath
	}
	// 模版渲染，并写入文件
	f, err := os.OpenFile(WFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		logrus.Errorf("[util - 模版文件打开] ｜ ❌ false |err ===》%s", err)
		return err, WFilePath
	}
	if err := tmpl.Execute(f, TplStruct); err != nil {
		logrus.Errorf("[util - 远程配置文件生成] ｜ ❌ false |err ===》%s ", err)
		return err, WFilePath
	}
	if DetailedOutput {
		logrus.Info("[util - 文件生成] ｜ ✅ true ")
	}
	// 模版渲染，并赋值给变量
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, TplStruct); err != nil {
		logrus.Errorf("[util - 远程配置文件变量赋值] ｜ ❌ false ｜err ===》%s ", err)
		return err, WFilePath
	}
	if DetailedOutput {
		logrus.Infof("[util - 文件生成 | 标准输出 |%s ]\n", WFilePath)
		// 模版渲染
		if err := tmpl.Execute(os.Stdout, TplStruct); err != nil {
			logrus.Errorf("[util - 远程配置文件输出到屏幕标准输出] ｜ ❌ false |err ===》%s |\n", err)
			return err, WFilePath
		}
	} else {
		// 模版渲染，并输出到屏幕标准输出
		if err := tmpl.Execute(os.Stdin, TplStruct); err != nil {
			logrus.Errorf("[util - 远程配置文件输出到屏幕标准输出] ｜ ❌ false |err ===》%s |\n", err)
			return err, WFilePath
		}
	}

	return nil, WFilePath
}

//func main()  {
//
//
//	type DataStruct struct {
//		ProjectName string
//		CMD         string
//		Program     string
//	}
//	// 实例化结构体
//	foo := DataStruct{
//		ProjectName: "AD-001",
//		Program:     "spring-gateway",
//		CMD:         "CMD11",
//	}
//	GenerateTpl("./data/shuchu.txt","./data/tpl.txt",&foo,false)
//}
