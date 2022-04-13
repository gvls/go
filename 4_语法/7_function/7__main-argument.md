## 2 argument

### 3  simple argument
Need `import "os"` 
order of argument must fixation
`os.Args[0]` is name of binary file
```go
os.Args		// It is []string
```


### 3  complex argument
Need `import "flag"` 
order of argument can `任意` 
All `argument` need add `-arg` from it

* get `string` 
```go
flag.StringVar(&stringVariateName, "arg", defaultValue, "explain for arg")

// arg : value after '-'
// defaultValue : if not get value for arg, value of stringVariateName is defaultValue

flag.Parse()	// get argument from `os.Args[1:]` 
```


* get `int` 
```go
flag.IntVar(&intVariateName, "arg", defaultValue, "explain for arg")

flag.Parse()
```
