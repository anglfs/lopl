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
��������������������������������
��Ȩ����������ΪCSDN�����������ڵ����ǡ���ԭ�����£���ѭ CC 4.0 BY-SA ��ȨЭ�飬ת���븽��ԭ�ĳ������Ӽ���������
ԭ�����ӣ�https://blog.csdn.net/hjxzb/article/details/84890959