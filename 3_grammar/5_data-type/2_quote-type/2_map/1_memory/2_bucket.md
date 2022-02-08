##  bucket
```shell
bucket 		_______________________
			|      |      |       |
			| hash |array |pointer|
			|      |      |       |
			-----------------------

hash    : high bit hash array
array   : save key-value(max is 8 key-value)
pointer : pointer next expand bucket

max count = 2^B
```

* array
```shell
array	_____________________________________
		|	  |     |     |     |     |     |
		| key | key | ... |value|value| ... |
		|     |     |     |     |     |     |
		-------------------------------------
go not use |key|value|key|value|...
by use this way : If length of key and value is differrent, it can needn't space to save padding(内存对齐)
```
