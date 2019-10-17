package main

import (
	"fmt"
)

package main

import (
   "fmt"
   "os"
   "os/signal"
   "syscall"

   "github.com/zhanben/go_site/app"
   "github.com/zhanben/go_site/tool/config"
   "github.com/zhanben/go_site/tool/db"
   "github.com/zhanben/go_site/tool/log"

   "go.uber.org/zap"
)

func main() {
   //Read config file
   err := config.ParseConfig()
   if err != nil {
      panic(fmt.Errorf("Failed to read config file: %s \n", err))
   }

   //Init log
   log.InitLog()

   //Init db connection
   db, err := db.InitDbConn()
   if err != nil {
      panic("connect db error!")
   }
   log.Logger.Info("Db init successful!")

   //start http sever
   startServer(db)
}
————————————————
版权声明：本文为CSDN博主「骑着炮弹进城」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/hjxzb/article/details/84890959