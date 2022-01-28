### 3  ++ 
`<variate>++` == `<variate> = <variate> + 1` 
only use by independence (`var a = b++` is error)

### 3  -- 
`<variate>--` == `<variate> = <variate> - 1` 
only use by independence (`var a = b--` is error)

### 3  + (`正号`) 
`+<number>`

### 3  - (`负号`) 
`-<number>`

### 3  + 
`1 + 2` 
If **all** variate is number, it is `加法运算`
If **all** variate is string, it is `拼接`

### 3  - 
`<number> - <number>` 

### 3  * 
`<value> * <value>`

### 3  / 
`<value> / <value>` 
If **all** number is `int` , the result is `int` 
If want result save `小数`, **all** `variate` must is `float...`
If want result save `小数` and **all** is `真实的数字不是变量` , `其中一个数字要带小数点`)

### 3  % (`取余`) 
`<number 1> % <number 2>` 
symbol of result is symbol of `number 1`
