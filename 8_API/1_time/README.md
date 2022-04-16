##  time

###   常量
秒之间的转换单位是1000 
相对大的秒不能通过除法得到小的秒 

* Nanosecond  == 1 `纳秒`
* Microsecond == 1 `微秒`
* Millisecond == 1 `毫秒`
* Second      == 1 `秒`
* Minute      == 1 `分钟`
* Hour        == 1 `小时`

###   休眠
`time.Sleep(常量)` 


###   获取时间
time.Now() 

* year            : `time.Now().Year()`
* month(alphabet) : `time.Now().month()`
* month(number)   : `int( time.Now().month() )`
* day             : `time.Now().Day()`
* hour            : `time.Now().Hour()`
* minute          : `time.Now().Minute()`
* second          : `time.Now().Second()`

###   获取格式化的时间
`fmt.Printf("...", time.Now(). ...)` 
`time.Now().Format("组合")` 

* year   : 2006	(6)
* month  : 01	(1)
* day    : 02	(2)
* hour   : 15	(3)
* minute : 04	(4)
* second : 05	(5)


###   unix时间戳
1970-1-1 到现在经过的秒
可以作为随机函数的种子
`time.Now().Unix() int64` 


###   unixnano时间戳
1970-1-1 到现在的纳秒，数据大于`int64` 的数据范围
可以作为随机函数的种子
`time.Now().UnixNano() int64` 
