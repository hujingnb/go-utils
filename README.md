# go-utils

[english](./README.en.md)

一些常用的函数集合

如果你遇到了类似的问题但在这里没有找到, 欢迎提交`issue`或`PR`, `PR`请提交到`dev`分支.

安装: `go get github.com/hujingnb/go-utils`

可在[godoc](https://pkg.go.dev/github.com/hujingnb/go-utils)中找到相关文档.

# 包说明

* `hstring`: [字符串操作](./hstring/README.md)
* `hhash`: [哈希函数](./hhash/README.md)
* `harray`: [数组操作](harray/README.md)
* `hmap`: [map操作](./hmap/README.md)
* `hstruct`: [结构体操作](hstruct/README.md)
* `hnumber`: [数字操作](hnumber/README.md)
* `hsystem`: [系统操作](hsystem/README.md)
* `hcommon`: [公共方法(如深度复制)](./hcommon/README.md)

# PR提交规则

1. 请求每次修改之后, 运行所有测试用例. 命令: `go test ./...`
2. 所有方法都要有对应的测试用例
3. 所有方法都要在`example_test.go`文件中给出示例
4. 所有对外访问的方法, 都要在`README`中标识
