package cmd

import (
	"aigc-go/internal/controller/getret"
	"aigc-go/internal/controller/getretbyuuid"
	"aigc-go/internal/controller/hello"
	"aigc-go/internal/controller/interceptor"
	"aigc-go/internal/controller/upload"
	uuid2 "aigc-go/internal/controller/uuid"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

// 自定义的拦截器，用于检查 JSON 响应中的 `code` 字段
func jsonResponseInterceptor(r *ghttp.Request) {
	// 继续处理请求，直到得到响应
	r.Middleware.Next()

	// 如果返回的是 JSON 格式，检查响应内容
	if r.Response.Status == 404 {
		r.Response.WriteStatus(404) // 设置状态码为 404
		err := r.Response.WriteTpl("404.html")
		if err != nil {
			return
		} // 输出自定义 404 页面
		r.Exit()
	}
}

var (
	Main = gcmd.Command{
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 创建一个HTTP服务器实例
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Use(jsonResponseInterceptor)
			// 路由分组
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareCORS)
				group.Bind(
					hello.NewV1(),
					upload.Uploadv1(),
					uuid2.Uuidv1(),
					getret.NewGetretV1(),
					getretbyuuid.NewGetretbyuuidV1(),
					interceptor.NewInterceptorV1(),
				)

			})
			// 启动服务器
			s.Run()
			return nil
		},
	}
)
