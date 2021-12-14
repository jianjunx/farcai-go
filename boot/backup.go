package boot

import (
	"context"
	"farcai-go/app/service"

	"github.com/adhocore/gronx/pkg/tasker"
	"github.com/gogf/gf/frame/g"
)

// 备份的定时任务
func backupClock() {
	if !g.Cfg().GetBool("database.backup") {
		return
	}
	// 开启定时任务
	taskr := tasker.New(tasker.Option{
		Verbose: true,
		Tz:      "Asia/Shanghai", // 时区
	})
	taskr.Task(g.Cfg().GetString("database.backupCron"), func(ctx context.Context) (int, error) {
		go service.Assets.BackupCategory()
		go service.Assets.BackupPost()
		go service.Assets.BackupUser()
		return 0, nil
	})
	// 在协程中执行定时任务
	go taskr.Run()
}
