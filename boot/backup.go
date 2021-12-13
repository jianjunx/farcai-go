package boot

import (
	"farcai-go/app/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
)

// 备份的定时任务
func backupClock() {
	if !g.Cfg().GetBool("database.backup") {
		return
	}
	// 开启定时任务
	_, err := gcron.Add(g.Cfg().GetString("database.backupCron"), func() {
		// 执行备份
		go service.Assets.BackupCategory()
		go service.Assets.BackupUser()
		go service.Assets.BackupPost()
	})
	if err != nil {
		g.Log().Error(err) // 输出log
	}
}
