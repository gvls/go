### 3  ++ 
like `<variate> = <variate> + 1` 
only use by independence (`var a = b++` is error)
`<variate>=1 <variate>++`

### 3  -- 
like `<variate> = <variate> - 1` 
only use by independence (`var a = b--` is error)
`<variate>=2 <variate>--`

### 3  + (`正号`) 
`+<number>`

### 3  - (`负号`) 
`-<number>`

### 3  * 
`<number> * <number>`

### 3  / 
`<number> / <number>` 
If all number is `int` , the result is `int` (remove `小数` )
If want result save `小数`, all `variate` must is `float...` (if all is `真实的数字不是变量` , `其中一个数字要带小数点`)

### 3  %(`取余`) 
`<number 1> % <number 2>` 
symbol of result is symbol of `number 1`

### 3  + 
`1 + 2` 
If **all** variate is number, it is `加法运算`
If **all** variate is string, it is `拼接`

### 3  - 
`<number> - <number>` 
