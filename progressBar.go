package meUtils

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"time"
)

func ProgressBar(FuncName string, Count int64) {
	count := Count
	// create and start new bar
	//bar := pb.StartNew(count)

	// start bar from 'default' template
	// bar := pb.Default.Start(count)

	// start bar from 'simple' template
	// bar := pb.Simple.Start(count)

	// start bar from 'full' template
	// bar := pb.Full.Start(count)

	tmpl := fmt.Sprintf("{{ red \"%s:\" }} {{ bar . \"[-\" \"-\" (cycle . \"—>\" \"↖\" \"↗\" \"↘\" \"↙\") \".\" \">]\"}} {{speed . | rndcolor }} {{percent .}}  {{string . \"my_green_string\" | green}} {{string . \"my_blue_string\" | blue}}", FuncName)
	// start bar based on our template
	bar := pb.ProgressBarTemplate(tmpl).Start64(count)
	// set values for string elements
	//bar.Set("my_green_string", "描述1").Set("my_blue_string", "描述2")
	for i := 0; i < int(count); i++ {
		bar.Increment()
		ProgressBarTask()
	}
	bar.Finish()
}

func ProgressBarTask() {
	time.Sleep(1 * time.Second)
	// 尽量不要有输出，否则会换行
	// Try not to have output, otherwise I will wrap.
	// Try not to have output, otherwise I will wrap.
	fmt.Println("ProgressBar test ..")
}

//需要指定进度条前缀、处理的次数，比如处理多少个文件，发出多少个请求之类的
//You need to specify the progress bar prefix, the number of times processed, such as how many files are processed, how many requests are made, and so on.
//总等待时间 = 次数 * 你程序消耗时间
//Total waiting time = times * your program consumes time
//func main() {
//	ProgressBar("test", 10)
//}
