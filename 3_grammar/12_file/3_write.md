## 2 write 
### 3  have `缓冲区` 
need `import ( "bufio" "os" )` 

* open file
```go
os.OpenFile(pathName string, flag int, perm FileMode) (f *File, err error)
```

* create `stream` 
```go
bufio.NewWriter(f) (w *Writer)
```

* read content

* write content from string to cache
```go
w.WriteString(str string)
```

* write content from cache to file
```go
w.Flush()
```

* close
```go
f.Close() error
```




### 3  read by once
Need `import ( "io/ioutil" )` 

```go
ioutil.ReadFile(pathName string) ([]byte, error)	// it open and close file automaticly
```

```go
ioutil.WriteFile(pathName string, data []byte perm os.FileMode) error
```


