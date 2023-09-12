package blog

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func loadEnv() {

	// 最好是 .env
	err := godotenv.Load(cfgFile)
	if err != nil {
		// 如果加载 .env 失败，打印 `'Error: xxx` 错误，并退出程序（退出码为 1）
		cobra.CheckErr(err)
	}

	// 打印 viper 当前使用的配置文件，方便 Debug.
	fmt.Fprintln(os.Stdout, "Using config file:", cfgFile)
}

func env(key, fallbackValue string) string {
	s, ok := os.LookupEnv(key)
	if !ok {
		return fallbackValue
	}
	return s
}
