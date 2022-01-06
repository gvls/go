## 2 `package`
`本质` : `创建不同的文件夹存放程序文件` 
`package` `封装` some `property` or `functiom` to let other method `调用` it
go manage file and directory structure of project by `package` 
Any source file belong a `package` 
One `package` include 0 ~ more file
Can't have `相同的函数名` or `全局变量` in `同一个package` 

* `directory` 
`package name` of current `directory`   ==   name of current `directory` 
They generally use lower case
they use `驼峰式` 

* file name
~~ `file name` of current file **generally** is same as `package name` ~~
use `_` replace `驼峰式` 


### 3  function of `package` 
distinguish `function` or `variate` which they name is same
Make manage project is easy
Control visit scope of `function` or `variate` 


### 3  `main` `pacakge` 
only `main` `package` can be compile to binary file
`main` function belong to `package main` 
