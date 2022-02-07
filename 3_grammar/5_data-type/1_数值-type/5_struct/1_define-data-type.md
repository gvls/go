##  define data type
define in `globle` 
```go
type <sturtName> struct { 	// if structName 大写首字母 : other package can visit
	<name> <type>			// field
	<name> <type>			// if name 大写首字母 : other package can visit 
	<name> <type>			// if name 大写首字母 : field can be 序列化
	[name, name type]		// define more field in one line
	...
}
```

###   point custom name for 序列化
Add `<反引号>json:"<new name>[,omitempty]" yaml:"<new name><反引号>` to end of `data type` 
Go will use `反射` to get `tag标签` 

* omitempty
If add `omitempty` (**can't** add ` ` to side of `,`), When do `序列化`, `property` will ignore `default value`
If want or nested `struct` be ignore in `序列化` , we need add `omitempty` and change it as `pointer` 

