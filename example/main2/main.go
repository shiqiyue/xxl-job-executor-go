package main

import (
	"fmt"
	xxl "github.com/shiqiyue/xxl-job-executor-go"
	"github.com/shiqiyue/xxl-job-executor-go/example/task"
	"log"
)

func main() {
	exec := xxl.NewExecutor(
		xxl.ServerAddr("http://192.168.3.132:8100"),
		xxl.AccessToken("fafab23afaffabdfhft324"), //请求令牌(默认为空)
		xxl.ExecutorIp("192.168.3.5"),             //可自动获取
		xxl.ExecutorPort("9998"),                  //默认9999（非必填）
		xxl.RegistryKey("test"),                   //执行器名称
		xxl.SetLogger(&logger{}),                  //自定义日志
	)
	exec.Init()
	//设置日志查看handler
	exec.LogHandler(func(req *xxl.LogReq) *xxl.LogRes {
		return &xxl.LogRes{Code: 200, Msg: "", Content: xxl.LogResContent{
			FromLineNum: req.FromLineNum,
			ToLineNum:   2,
			LogContent:  "这个是自定义日志handler",
			IsEnd:       true,
		}}
	})
	//注册任务handler
	exec.RegTask("task.test", task.Test)
	exec.RegTask("task.test2", task.Test2)
	exec.RegTask("task.panic", task.Panic)
	log.Fatal(exec.Run())
}

//xxl.Logger接口实现
type logger struct{}

func (l *logger) Info(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf("自定义日志 - "+format, a...))
}

func (l *logger) Error(format string, a ...interface{}) {
	log.Println(fmt.Sprintf("自定义日志 - "+format, a...))
}
