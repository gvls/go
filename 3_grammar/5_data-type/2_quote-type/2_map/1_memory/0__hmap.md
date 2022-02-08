##  hmap
a header for a go map
```shell
type hmap struct{				// implement 散列表/hashTable. use 拉链法 解决碰撞
	count 	   int				// count of key-value
	flags 	   uint8			// if writing or 迁移
	B	  	   uint8			// relation to expand capacity. == bit count of low bit hash
	noverflow  uint16			// 近似数(when B>=16) of o_bucket 
	hash0	   uint32			// random hash seed(prevent 哈希碰撞). When map create, it will be create by fastrand()

	buckets	   unsafe.Pointer	// 哈希桶 point to [...]bucket.
	oldbuckets unsafe.Pointer	
	nevacuate  uintptr			// process of expand cap. address which less than it of bucket already 迁移

	mapextra   *mapextra		// reduce GC scan
}
```
