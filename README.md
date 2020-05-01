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

import (
	"fmt"
	"github.com/gscsdlz/json_reader"
	"os"
)

func main() {
	str := `
{
  "obj": {
    "k1": 1000,
    "k2": "2"
},
  "arr": [
    "v1",
    {
      "obj2": {
        "k1": 1, 
        "k2": "hello"
      }
    }
  ]
}
`
	r, err := json_reader.Parse([]byte(str))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(r.Get("obj").Get("k1").Number())
	arrItem1 := r.Get("arr").Get(1)
	fmt.Println(arrItem1.Get("obj2").Get("k2").String())
}

```