slice初始化

```go
s :=make([]int,len,cap) // 通过内置函数 make() 初始化切片s，[]int 标识为其元素类型为 int 的切片
s :=[]int {1,2,3 }  // 直接初始化，[]表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。
s := arr[:] // 初始化切片 s，是数组 arr 的引用
s := arr[startIndex:endIndex] // 将arr中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片
s := arr[startIndex:] // 默认 endIndex 时将表示一直到arr的最后一个元素
s := arr[:endIndex] // 默认 startIndex 时将表示从 arr 的第一个元素开始
```

