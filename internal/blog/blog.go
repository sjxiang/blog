package blog

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"github.com/sjxiang/blog/pkg/zop"
)

// NewBlogCommand 创建一个 *cobra.Command 对象. 之后，可以使用 Command 对象的 Execute 方法来启动应用程序.
func NewBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "博客",
		// 命令的简短描述
		Short: "Go 实战项目",
		// 命令的详细描述
		Long: `详情：https://github.com/sjxiang/blog`,

		// 命令出错时，不打印帮助信息。不需要打印帮助信息，设置为 true 可以保持命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {

			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	// 调用链路：Execute -> init -> Run
	cobra.OnInitialize(loadEnv)

	// 选项，e.g. --config=xxx
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./configs/.env", "miniblog 的配置文件路径，若字符串为空，则为无配置文件。")

	return cmd
}

var cfgFile string

// run 函数是实际的业务代码入口函数.
func run() error {

	// 优雅关停
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	var (
		logDisableCaller, _     = strconv.ParseBool(env("LOG_DISABLE_CALLER", ""))
		logDisableStacktrace, _ = strconv.ParseBool(env("LOG_DISABLE_STACKTRACE", ""))
		logLevel                = env("LOG_LEVEL", "")
		logFormat               = env("LOG_FORMAT", "")
		logOutputPaths          = env("LOG_OUTPUT_PATHS", "")

		runMode                 = env("RUN_MODE", "")
		port, _                 = strconv.Atoi(env("PORT", ""))
	)

	// var (
	// 	dbUser = env("DB_USER", "")
	// 	dbPass = env("DN_PASS", "")
	// 	dbHost = env("DB_HOST", "")
	// 	dbPort = env("DB_PORT", "")	
	// 	dbName = env("DB_NAME", "")	
	// 	dbMaxIdleConnection, _     = strconv.Atoi(env("DB_MAX_IDLE_CONNECTION", ""))
	// 	dbMaxOpenConnection, _     = strconv.Atoi(env("DB_MAX_OPEN_CONNECTION", ""))
	// 	dbMaxConnectionLifeTime, _ = strconv.ParseFloat(env("DB_MAX_CONNECTION_LIFE_TIME", ""),  64)
	// 	dbLogLevel, _              = strconv.Atoi(env("DB_LOG_LEVEL", ""))
	// )
	
	// 初始化日志
	zop.Init(logOptions(logDisableCaller, logDisableStacktrace, logLevel, logFormat, logOutputPaths))
	defer zop.Sync()  

	// 设置 Gin 模式
	gin.SetMode(runMode)

	r := gin.New()

	r.NoRoute(func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.RequestURI, "/v1") || strings.HasPrefix(ctx.Request.RequestURI, "/api") {
			// controller.RelayNotFound(c)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code":    10003, 
			"message": "Page not found", 
		})
	})

	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// 构建 server
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

		
	go func() {
		<-ctx.Done()
		zop.Infow("shutting down server")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		err := httpSrv.Shutdown(ctx)
		if err != nil {
			zop.Errorw("shutdown server", "err", err)
			os.Exit(1)
		}
	}()

	zop.Infow("starting server", "addr", httpSrv.Addr)

	 
	err := httpSrv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		// zop.Fatal("listen and serve: %w", err)  // 直接崩了，啥细节也看不到
		return fmt.Errorf("listen and serve: %w", err)
	}
	
	return nil
}

