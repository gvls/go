##  visit element
format is like array

###   visit element by byte (can't visit china)
```shell
value : "520中"

physics : [7]byte{53,50,48, 228,184,173}

        5   2   0       中
char    |   |   |    /  |  \ 
      _________________________
      |   |   |   |   |   |   |
value | 53| 50| 48|228|184|173|
      |   |   |   |   |   |   |
      -------------------------
index   0   1   2   3   4   5		cell unit : byte
```
* direct
```go
<variate>[<index>] 	// index is begin from 0
```
* for




###   visit element by character (can visit china)
```shell
value : "520中"

physics : [4]rune{53,50,48, 20013}

           5         2         0          中
char       |         |         |          |   
      ___________________________________________
      |         |         |         |           |
value |    53   |    50   |    48   |   20013   |
      |         |         |         |           |
      -------------------------------------------
index      0         1         2          3    		cell unit : rune
```
* transform to []rune

* range
china character will be compound to one element
