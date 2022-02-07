##  map
It is a `key-value` data structure
It's `存储` is `无序` 
Capacity of `map` is dynamic increase (not replace old memory immediately. old memory be free by `gc`)
Element of `map` can't be `取地址` 

```shell
type hmap struct{				// implement 散列表. use 拉链法 解决碰撞
	count 	   int				// count of key-value
	flags 	   uint8
	B	  	   uint8
	noverflow  uint16
	hash0	   uint32			// hash seed

	buckets	   unsafe.Pointer	// point to [...]bucket.
	oldbuckets unsafe.Pointer	
	nevacuate  uintptr			//

	extra	   *mapextra		//
}
```
```shell
buckets		 _____________________________
			 |      |      |      |      |
head node	 |bucket|bucket|bucket|bucket|
			 |      |      |      |      |
			 -----------------------------
				|      |      |      |
				v 	   v      v      v
link table	 bucket  bucket bucket bucket
				|      |      |      |
				v      v      v      v
		   	   ...    ...    ...    ...
```

* loadFactor 加载因子
loadFactor =  散列包含的元素数 / 位置总数
If loadFactor more high, 冲突产生概率 more high
When loadFactor 大于阈值, len(map)/2^B > 6.5, visit efficiency change to low and Go will expand cap for map 
when expand cap, it point to old buckets. when next time we visit old value, copy value of oldbuckets to buckets and cancel quote for old value. memory of oldbuckets be free by gc

* 哈希函数
哈希值 = 哈希fuc(key)
result is lowValue and highValue
lowValue is local place of buckets
hightValue is local key (node of bucket link table)
