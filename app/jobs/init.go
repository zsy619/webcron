package jobs

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"haedu.gov.cn/dzz/tt/app/models"
)

func InitJobs() {
	list, _ := models.TaskGetList(1, 1000000, "status", 1)
	for _, task := range list {
		job, err := NewJobFromTask(task)
		if err != nil {
			logs.Error("InitJobs:", err.Error())
			continue
		}
		AddJob(task.CronSpec, job)
	}
}

func runCmdWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		logs.Warn(fmt.Sprintf("任务执行时间超过%d秒，进程将被强制杀掉: %d", int(timeout/time.Second), cmd.Process.Pid))
		go func() {
			<-done // 读出上面的goroutine数据，避免阻塞导致无法退出
		}()
		if err = cmd.Process.Kill(); err != nil {
			logs.Error(fmt.Sprintf("进程无法杀掉: %d, 错误信息: %s", cmd.Process.Pid, err))
		}
		return err, true
	case err = <-done:
		return err, false
	}
}
