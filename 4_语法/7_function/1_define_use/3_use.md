## 2 use
```go
// If want to ignore some `return value` use `_` to `占位` 
[<variate>[, ...] = ]   <name>([<parameter>[, ...]>)
```
* Function will get the value of parameter
If value of parameter is `值` (`值类型`), it like copy value of `parameter` to `variate` of `function` 
If value of parameter is `address` (`引用类型` or `&值类型`), it like bind `parameter` to `variate` of `function` 
