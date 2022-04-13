## 2 `recover` 
It can catch error and let procedure **not exist**

```go
err := recover()	// catch error which `之前的代码发生` 
if err != nil {
			// here handle error
}
```

```go
defer func () {		// 匿名函数
	// here catch and handle error
}()
```


