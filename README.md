# JSON_READER
> Warning：实验性模块，不能用于生产环境，不提供维修服务哈

## 目的
Golang中如果要获取JSON字段中某个字段的key需要提前绑定struct或者用interface接收，
但是interface需要不断的类型推导，代码冗长。<br/>
JSON_READER支持使用链式调用取出JSON中的任意value

## 用法
### func Parse(data []byte) (JsonReader, error)
解析JSON并返回JsonReader对象

### func (r *JsonReader) Get(key ...interface{}) JsonReader
链式调用获取对象或者数组的元素，key支持string或者int

### func (r *JsonReader) Number() float64
返回当前number类型的值

### func (r *JsonReader) String() string
返回当前string类型的值

### func (r *JsonReader) Len() int
返回当前数组或者对象的长度，用于支持遍历

## 示例
```go
package main

import (
	"fmt"
	"github.com/gscsdlz/json_reader"
	"os"
)

func main() {
	str := `{"obj":{"k1":1000,"k2":"2"},"arr":[{"k1":1,"k2":"hello1"},{"k1":2,"k2":"hello2"}]}`
	r, err := json_reader.Parse([]byte(str))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(r.Get("obj").Get("k1").Number())

	arrItem := r.Get("arr")
    for i := 0; i < arrItem.Len(); i++ {
	    fmt.Println(arrItem.Get(i).Get("k1").Number())
	    fmt.Println(arrItem.Get(i).Get("k2").String())
    }
}

```

输入原始JSON
```json
{
    "obj":{
        "k1":1000,
        "k2":"2"
    },
    "arr":[
        {
            "k1":1,
            "k2":"hello1"
        },
        {
            "k1":2,
            "k2":"hello2"
        }
    ]
}
```

输出
```text
1000
1
hello1
2
hello2
```