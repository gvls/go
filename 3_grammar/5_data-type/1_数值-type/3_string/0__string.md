##  string
If have `"` in `"<string>"` , `"` need write `\"`
Go use UTF-8 to `标识` Unicode text
底层 like **read only** slice([]byte) without cap : it can be slice 

* basic data struct
```shell
type StringHeader struct {
	Data uintptr	// point to actual data(it will be reusing by many variate)
	Len int			// length of value
}
```

* value which Data point to 
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
