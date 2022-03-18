##  map
It is a `key-value` data structure
It's `存储` is `无序` 
Capacity of `map` is dynamic increase (not replace old memory immediately. old memory be free by `gc`)
Element of `map` can't be `取地址` 
**not** 线程安全


###   make map 线程安全
* add read write lock
lock total map
suit : a few 并发

* concurrent-map 库
lock part map
suit : maintain 映射表

* use sync.map
空间换时间
suit : count of read > count of write   and write most is insert map
