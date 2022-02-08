##  loadFactor
加载因子(阈值、判断是否扩容)
balance 冲突产生概率 和space use ratio
If loadFactor more high, 冲突产生概率 more high
If loadFactor more low,  space use ratio more low
折中办法   : loadFactor =  total key-value / total bucket(include obucket)
**Go折中办法** : loadFactor = length of buckets / 2^B(count of bucket)
