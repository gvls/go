##  slice
动态数组 : a dynamic window( [indexi, indexj) ) which point to array
It's **length can be automatic and dynamic change**(when capacity > length) in procedure running
It is a **struct** and quote type and basic on **array** 
`slice` can quote more than one `array` 
Use it is like `array` 
capacity >= length

```shell
type slice struct {			
	array unsafe.Pointer 	// point to first element of array
	len   int 				// count of element of slice
	cap   int 				// length of capacity (may != length of array)
}

logic : s := make([]int, 2, 4);  s[0]=10;  s[1]=30

physics :

   		_________________________
		|       |       |       |
slice	|pointer|   2   |   4   |
		|       |       |       |
		-------------------------
		    |
			v
        _________________________________
		|       |       |       |       |
array	|  10   |  20   |   0   |   0   |
		|       |       |       |       |
		---------------------------------


```
