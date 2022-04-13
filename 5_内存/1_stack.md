## 2 `stack` 
```shell
		 func		 	 func 
|    |	| |  |	|    |	| ^  |	|    |	<- free space
|    |	| v  |	|    |	| |  |	|    |
|    |	|    |	|    |	|    |	|    |
|    |	|    |	|func|	|    |	|    |
|func|	|func|	|func|	|func|	|func|	<- use space
|main|	|main|	|main|	|main|	|main|
```
generally is save `basic data type `
When `function` is running , `stack` will get a part of space for this `function` 
When `function` is end, `stack` will recycle space which `function` use


