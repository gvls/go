## 2 number
number in the end of `data type` is bit
1 byte == 8 bit

### 3  int
#### 4   have symbol
First bit be used to symbol

* `int` (default)
If 32 bit system: four byte
-2^31 ~ -2^31-1

If 64 bit system: eight byte
-2^63 ~ -2^63-1

* `int8` one byte
-128 ~ 127(2^7 -1)

* `int16` two byte
-2^15 ~ 2^15-1

* `int32 , rune` four byte
visit element by byte (can't visit china character)
-2^31 ~ 2^31-1
`rune` is use in Unicode

* `int64` eight byte
-2^63 ~ 2^63-1


#### 4   haven't symbol
* `uint` 
If 32 bit system: four byte
0 ~ 2^32-1

If 64 bit system: eight byte
0 ~ 2^64-1

* `uint8 , byte` one byte
0 ~ 255 (2^8 -1)

* `uint16` two byte
0 ~ 2^16-1

* `uint32` four byte
0 ~ 2^32-1

* `uint64` eight byte
0 ~ 2^64-1

* `uintptr` 
save pointer



### 3  float
float = symbol + `指数` + `尾数`
Some value may have `精度损失` and `float64` lose less than `float32` 

* `float32` four byte
`-3.403E38` ~ `3.403E38`

* `float64` eight byte (default, recommend)
`-1.798E308` ~ `1.798E308`

#### 4   evaluation
* `十进制` 
such as: `10.0` , `.123`

* Scientific notation
such as: `10.1243e2 (10.12*10^2)` , `10.1243E2` , `10.1243e-2` 



### 3  complex
* `complex64` four byte
`实数` , `虚数` 

* `complex128` eight byte
`实数` , `虚数` 
