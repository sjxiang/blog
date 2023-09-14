package blog

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"github.com/sjxiang/blog/internal/blog/store"
	"github.com/sjxiang/blog/pkg/db"
	"github.com/sjxiang/blog/pkg/middleware"
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
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./config/.env", "miniblog 的配置文件路径，若字符串为空，则为无配置文件。")

	return cmd
}

var cfgFile string

// run 函数是实际的业务代码入口函数.
func run() error {

	// 初始化环境变量 env

	var (
		logDisableCaller, _     = strconv.ParseBool(env("LOG_DISABLE_CALLER", ""))
		logDisableStacktrace, _ = strconv.ParseBool(env("LOG_DISABLE_STACKTRACE", ""))
		logLevel                = env("LOG_LEVEL", "")
		logFormat               = env("LOG_FORMAT", "")
		logOutputPaths          = env("LOG_OUTPUT_PATHS", "")

		runMode                 = env("RUN_MODE", "")
		port, _                 = strconv.Atoi(env("PORT", ""))
	
		dbDataSource               = env("DB_DATA_SOURCE", "")
		dbMaxIdleConnection, _     = strconv.Atoi(env("DB_MAX_IDLE_CONNECTION", ""))
		dbMaxOpenConnection, _     = strconv.Atoi(env("DB_MAX_OPEN_CONNECTION", ""))
		dbMaxConnectionLifeTime, _ = strconv.Atoi(env("DB_MAX_CONNECTION_LIFE_TIME", ""))
		dbLogLevel, _              = strconv.Atoi(env("DB_LOG_LEVEL", ""))
	)
	
	// 初始化日志 log
	zop.Init(logOptions(logDisableCaller, logDisableStacktrace, logLevel, logFormat, logOutputPaths))
	defer zop.Sync()  

	// 初始化数据库 db
	dbOptions := mysqlOptions(dbDataSource, dbMaxIdleConnection, dbMaxOpenConnection, dbMaxConnectionLifeTime, dbLogLevel)
	db, err := db.NewMySQL(dbOptions)
	if err != nil {
		return err
	}
	
	// 构建 store（包变量有点恶心）
	store := store.NewStore(db)
	
	// 设置 Gin 模式
	gin.SetMode(runMode)

	r := gin.New()
	
	// 注册全局中间件
	mws := []gin.HandlerFunc{gin.Recovery(), middleware.RequestID(), middleware.CorsV1()}
	r.Use(mws...)

	// 注册路由
	if err := setupRoute(store, r); err != nil {
		return err
	}
	
	// 构建 server
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

		
	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zop.Fatalw("listen and serve: %w", err) 
		}
	}()

	// 方便排障
	zop.Infow("starting server", "addr", httpSrv.Addr)

    quit := make(chan os.Signal, 1)
    // kill 默认会发送 syscall.SIGTERM 信号
    // kill -2 发送 syscall.SIGINT 信号，我们常用的 CTRL + C 就是触发系统 SIGINT 信号
    // kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
    <-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
    zop.Infow("Shutting down server ...")

	// 10s，结束扫尾工作
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	// 通知 goroutine，风紧扯呼
    if err := httpSrv.Shutdown(ctx); err != nil {
		zop.Errorw("shutdown server", "err", err)
        return err
    }
	

	return nil
}

