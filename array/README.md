数组处理 

导入: `import "github.com/hujingnb/go-utils/harray"`

主要方法如下: 

* `Chunk`:  将数组分割成多个
* `Shuffle`: 随机打乱数组顺序
* `InArray`: 检查元素是否在数组中
* `IndexOf`: 获取元素在数组中首次出现的索引
* `Unique`: 数组去重
* `Equal`: 判断两数组是否相等
* `Intersect`: 取数组交集
* `Diff`: 取数组差集
* `BinarySearch`: 二分查找
* 排序操作: 
  * `SortBubble`: 冒泡排序
  * `SortInsert`: 插入排序
  * `SortSelect`: 选择排序
  * `SortQuick`: 快排
  * `SortMerge`: 归并排序

示例代码见: [example_test.go](./example_test.go)