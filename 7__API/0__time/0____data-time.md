## 2 data time
`import "time"` 

### 3  `time constant`
`time.<cost>` 
`秒之间的转换单位是1000` 
`相对大的秒不能通过除法得到小的秒` 

* Nanosecond  == 1 `纳秒`
* Microsecond == 1 `微秒`
* Millisecond == 1 `毫秒`
* Second      == 1 `秒`
* Minute      == 1 `分钟`
* Hour        == 1 `小时`

### 3  sleep
`time.Sleep(<statement of time constant>)` 


### 3  get current local date and time
`time.Now()` 

* year            : `time.Now().Year()`
* month(alphabet) : `time.Now().month()`
* month(number)   : `int( time.Now().month() )`
* day             : `time.Now().Day()`
* hour            : `time.Now().Hour()`
* minute          : `time.Now().Minute()`
* second          : `time.Now().Second()`

### 3  get current local date and time by format
`fmt.Printf("...", time.Now(). ...)` 
`time.Now().Format("<组合>")` 

* year   : 2006	(6)
* month  : 01	(1)
* day    : 02	(2)
* hour   : 15	(3)
* minute : 04	(4)
* second : 05	(5)


### 3  `unix时间cuo` 
It is second which from 1970-1-1 to current
It be use to get random seed
`time.Now().Unix() int64` 


### 3  `unixnano时间cuo` 
It is nanosecond which from 1970-1-1 to current and it is bigger than `int64` 
It be use to get random seed
`比unix更加随机` 
`time.Now().UnixNano() int64` 
