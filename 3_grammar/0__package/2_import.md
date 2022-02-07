##  import package
Write it after `package name` 
If `package` is `内置标准库`, we needn't write `path` 
If `package` is custom, we need write `绝对路径` from `GOPATN/src/` 
When import `package` , we must use it

* one line
`import "[path]<name>[/<name>]"` 

* multi-line
```go
import (
	"[path]<name>[/<name>]"
	"[path]<name>[/<name>]"
	<new name> "[path]<name>[/<name>]"	// alias new name and old name can't be use 
	_ "[path]<name>[/<name>]"			// not import this package
)
```

