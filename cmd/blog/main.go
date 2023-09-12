package main

import (
	"os"

	_ "go.uber.org/automaxprocs"

	"github.com/sjxiang/blog/internal/blog"
)

func main() {
	// 挺鸡肋的
	command := blog.NewBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
