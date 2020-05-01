# JSON_READER
> Warning：实验性模块，不能用于生产环境，不提供维修服务哈

## 目的
Golang中如果要获取JSON字段中某个字段的key需要提前绑定struct或者用interface接收，
但是interface需要不断的类型推导，代码冗长。<br/>
JSON_READER支持使用链式调用取出JSON中的任意value

## 用法
### Parse(data []byte) (JsonReader, error)
解析JSON并返回JsonReader对象

### Get(key ...interface{}) JsonReader
链式调用获取对象或者数组的元素，key支持string或者int

### Number() float64
返回当前number类型的值

### String() string
返回当前string类型的值

## 示例
```go
package main
import "github.com/gscsdlz/json_reader"

func main() {
}
```