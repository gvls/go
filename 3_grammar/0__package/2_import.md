## 2 import `package`
Write it after `package name` 
If `package` is `内置标准库`, we needn't write `path` 
If `package` is custom, we need write `绝对路径` from `GOPATN/src/` 
If import `package` and not use, build will error

### 3  one line
`import "[path]<name>[/<name>]"` 

### 3  multi-line
```go
import (
	"[path]<name>[/<name>]"
	"[path]<name>[/<name>]"
)
```

### 3  alias
When define new name for `package`, we can't use old name of `package` 

```go
import (
	<new name> "<name>[/<name>]"
	...
)
```
