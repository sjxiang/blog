package blog

import (
	"fmt"

	"github.com/spf13/cobra"
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
	cobra.OnInitialize()

	// 选项，e.g. --config=xxx
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./configs/dev.yaml", "miniblog 的配置文件路径，若字符串为空，则为无配置文件。")

	return cmd
}

var cfgFile string

// run 函数是实际的业务代码入口函数.
func run() error {
	fmt.Println(cfgFile)
	return nil
}
