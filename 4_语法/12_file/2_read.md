## 2 `read` 
### 3  have `缓冲区` 
Need `import ( "bufio" "os" )` 
suit big file

* open
```go
os.Open(pathName string) (f *File, err error)	// f is 对象，指针，文件句柄
```

* create `stream` 
```go
bufio.NewReader(f *File) (r *Reader)
```

* read content
If `err` == `io.EOF` , already read to end of file
```go
r.ReadString(<这次读取的末尾字符>) (str string, err error)
```

* close
```go
f.Close() error
```

### 3  read by once
Need `import ( "io/ioutil" )` 
suit small file


```go
ioutil.ReadFile(pathName string) ([]byte, error)	// it open and close file automaticly
```


