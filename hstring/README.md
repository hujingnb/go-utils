字符串处理 

导入: `import "github.com/hujingnb/go-utils/hstring"`

主要方法如下: 

* `Reverse`: 字符串反转
* `CamelToSnake`: 驼峰转下划线
* `SnakeToCamel`: 下划线转驼峰
* `ToString`: 将变量转为字符串
* `Pad`: 将字符串填充到指定长度
* `Join`: 用字符串拼接数组元素
* 哈希算法
  * `Get32Md5`: 获取32位 md5值
  * `Get16Md5`: 获取16位 md5值
  * `GetSha1`: sha1加密

示例代码见: [example_test.go](./example_test.go)