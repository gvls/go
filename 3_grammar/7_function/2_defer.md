## 2 `defer` | `延时执行`  
When `function` run to end, `及时` release resource

### 3  format
```go
defer <statement>
```


### 3  process
If `function` define `variate` for `return list` , `defer` may change value of it

1. procedure run into one `function` 
2. run to line of `defer` 
3. make `statement` of `defer` into independent `栈` and `variate` which it include will be copy to independent `栈` 
4. execute `statement` of `return` and **evaluation** value of `statement` to temporary `variate` or `variate` of `return list` 
5. **execute** `statement` of `defer` of independent `栈` by `后进先出` 
6. **return** `return value` and finish `function` 
