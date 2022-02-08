##  hash function
```shell
"high bit hash"+"low bit hash" = hashFuc(key)
low bit hash   : local key belong which bucket
hight bit hash : local key belong which cell in bucket

if cpu support aes, Go will use aes hash otherwise use memhash
```

* 哈希函数
哈希值 = 哈希fuc(key)
result is lowValue and highValue
lowValue is local place of buckets
hightValue is local key (node of bucket link table)
