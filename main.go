package main

import (
	"aigc-go/internal/cmd"
	_ "aigc-go/internal/logic"
	_ "aigc-go/internal/packed"
	"aigc-go/utility"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "path/filepath"
)

func init() {
	utility.Handle2api()
}
func main() {
	cmd.Main.Run(gctx.GetInitCtx())

}
