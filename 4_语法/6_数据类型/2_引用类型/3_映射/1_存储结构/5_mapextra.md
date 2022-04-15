##  mapextra

* 背景
当 map 的 key 和 value 都是指针 并且他们的长度都没有超过 128 时(当 key 或 value 的长度超过 128 时, go 在 map 中会使用指针存储), 该 map 的 bucket 类型会被标注为不含有指针, 这样 gc 不会扫描该 map。

* 问题
bucket 的底层结构 bmap 中含有一个指向溢出桶的指针(uintptr类型, uintptr指针指向的内存不保证不会被 gc free 掉), 当 gc 不扫描该结构时, 该指针指向的内存会被 gc free 掉

因此在 hmap 结构中增加了 mapextra 字段, 其中 overflow 是一个指向保存了所有 hmap.buckets 的溢出桶地址的 slice 的指针
相对应的 oldoverflow 是指向保存了所有 hmap.oldbuckets 的溢出桶地址的 slice 的指针

只有当 map 的 key 和 elem 都不含指针时这两个字段才有效, 因为这两个字段设置的目的就是避免当 map 被 gc 跳过扫描带来的引用内存被 free 的问题, 当 map 的 key 和 elem 含有指针时, gc 会扫描 map, 从而也会获知 bmap 中指针指向的内存是被引用的, 因此不会释放对应的内存

```shell
tyep mapextra struct {
	overflow     *[]*bmap
	oldoverflow  *[]*bmap
	nextOverflow *bmap
}
```
