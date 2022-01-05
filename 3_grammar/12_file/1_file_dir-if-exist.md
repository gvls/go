## 2 file or directory if exist
Need `import "os"` 

1. If `err` is `nil`, file exist
2. If use `os.IsNotExist(err)` to judge `err` is `true`, file not exist
3. If `err` is other type, can't judge if(`是否`) exist
```go
os.Stat(pathName string) (fi FileInfo, err error)
```


