# go-utils

[english](./README.en.md)

一些常用的函数集合

如果你遇到了类似的问题但在这里没有找到, 欢迎提交`issue`或`PR`, `PR`请提交到`dev`分支.

安装: `go get github.com/hujingnb/go-utils`

可在[godoc](https://pkg.go.dev/github.com/hujingnb/go-utils)中找到相关文档.

# 包说明

## `hstring`

字符串相关操作. [Example](./hstring/example_test.go)

导入: `import "github.com/hujingnb/go-utils/hstring"`

方法: 

* `Reverse`: 字符串反转
* `CamelToSnake`: 驼峰转下划线
* `SnakeToCamel`: 下划线转驼峰
* `ToString`: 将变量转为字符串
* `Pad`: 将字符串填充到指定长度
* `Join`: 用字符串拼接数组元素

## `hhash`

哈希函数. [Example](./hhash/example_test.go)

导入: `import "github.com/hujingnb/go-utils/hhash"`

方法:

* `Md5By32`: 获取32位 md5值
* `Md5By16`: 获取16位 md5值
* `Sha1`: sha1加密
* `Sha256`
* `Sha512`
* `Crc32`

## `harray`

数组操作. [Example](./harray/example_test.go)

导入: `import "github.com/hujingnb/go-utils/harray"`

方法:

* `Chunk`:  将数组分割成多个
* `Shuffle`: 随机打乱数组顺序
* `InArray`: 检查元素是否在数组中
* `IndexOf`: 获取元素在数组中首次出现的索引
* `Unique`: 数组去重
* `Equal`: 判断两数组是否相等
* `EqualIgnoreOrder`: 判断两数组是否相等, 无视顺序
* `Intersect`: 取数组交集
* `Diff`: 取数组差集
* `BinarySearch`: 二分查找
* `GetSureRandArr`: 使用相同的种子生成相同随机数组
* `Count`: 统计数组中的元素数量
* `Reverse`: 反转数组
* 排序操作:
    * `SortBubble`: 冒泡排序
    * `SortInsert`: 插入排序
    * `SortSelect`: 选择排序
    * `SortQuick`: 快排
    * `SortMerge`: 归并排序

## `hmap`

map 操作. [Example](./hmap/example_test.go)

导入: `import "github.com/hujingnb/go-utils/hmap"`

方法:

* `IterateByOrder`: 字典序遍历 map
* `Equal`: 比较两 map 是否相等

## `hstruct`

结构体操作. [Example](./hstruct/example_test.go)

导入: `import "github.com/hujingnb/go-utils/hstruck"`

方法:

* `ToMap`: 转 map
* `Name`: 获取结构体的名称

## `hnumber`

数字操作. [Example](./hnumber/example_test.go)

导入: `import "github.com/hujingnb/go-utils/hnumber"`

方法:

* `HexInput`: 按指定进制读取字符串
* `HexOutput`: 将数字转为目标进制
* `RandRange`: 指定范围随机一个整数

## `hsystem`

系统级操作. [Example](./hsystem/example_test.go)

导入: `import "github.com/hujingnb/go-utils/hsystem"`

方法:

* `PrintOutAndErrToFile`: 将标准输出和错误输出写入到文件

## `hcommon`

一些公共方法. [Example](./hcommon/example_test.go)

导入: `import "github.com/hujingnb/go-utils/hcommon"`

方法:

* `Copy`: 对任意类型变量进行深度复制

# PR提交规则

1. 请求每次修改之后, 运行所有测试用例. 命令: `go test ./...`
2. 所有方法都要有对应的测试用例
3. 所有方法都要在`example_test.go`文件中给出示例
4. 所有对外访问的方法, 都要在`README`中标识
