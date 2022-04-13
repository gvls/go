##  bucket
* 编译前
```shell
type bmap struct{
	tophash [bucketCnt]uint8 // 存储 key 的哈希的高8位
	// 有些字段在运行后，通过地址方式访问
}
```

* 编译后
```go
type bmap struct{
	topbits		[8]uint8 		// 用于确认 key-value 在那个位置
	keys 		[8]keytype   	// 哈希函数计算得到的 低8位 一样
	values 		[8]valuetype
	pad 		uintptr 		// 使用keys/values而不是key/value/key/value/...节省字节对齐带来的空间浪费，也就是可以忽略该字段
	overflow 	uintptr 		// 指向 溢出桶。当有新的key-value插入，并且该桶已经满，就使用溢出桶存储
}
```
