##  for-range
可以遍历 `string` 、`array` 、 `slice` 、 `channel` 、 `map` 

* range 前面的变量
多次循环共用同一个变量
在出现`range` 后，直接**复制**原数据进行遍历。（若在`for-range` 里面修改原数据的值，`value`依然是旧数据的值，但是类似原数据是一个切片，也就是保存数组的地址，那么`value` 会被修改为新数据）



###   格式
* string
```go
var str = "hello world"
for 索引, 值 := range str { // 索引的数据类型是int，值的数据类型是rune(int32)，可以正常获取到中文
	...
}
```
* slice
```go
var s = []int{1, 2, 3}
for 索引, 值 := range s { // 在进入前，循环的次数已经被固定，在里面添加的元素 不会被访问到。索引的数据类型是int
	...
}
```
* map
```go
var ma = make(map[int]string)
for key, value := range ma { // 在循环过程中，若添加元素，循环的次数会变化。key-value 是无序获取的
	...
}
```
* channel
```go
var c = make(chan int, 10)
c <- 1
close(c)
for 值 := range c {
	...
}
```

