# ==============================================================================
# 定义全局 Makefile 变量方便后面引用

COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
# 项目根目录
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/ && pwd -P))
# 构建产物、临时文件存放目录
OUTPUT_DIR := $(ROOT_DIR)/_output

# ==============================================================================
# 定义 Makefile all 伪目标，执行 `make` 时，会默认会执行 all 伪目标
.PHONY: all
all: format build

# ==============================================================================
# 定义其他需要的伪目标

.PHONY: build
build: tidy # 编译源码，依赖 tidy 目标自动添加/移除依赖包.
	@go build -v -o $(OUTPUT_DIR)/blog $(ROOT_DIR)/cmd/blog/main.go

.PHONY: format
format: # 格式化 Go 源码.
	@gofmt -s -w ./
        
.PHONY: tidy
tidy: # 自动添加/移除依赖包.
	@go mod tidy
        
.PHONY: clean
clean: # 清理构建产物、临时文件等.
	@-rm -vrf $(OUTPUT_DIR)


.PHONY: print
print: # 打印变量
	@echo $(OUTPUT_DIR)
	@echo $(ROOT_DIR)
        

# tips
# 空格和 tab，有误区，真的会谢