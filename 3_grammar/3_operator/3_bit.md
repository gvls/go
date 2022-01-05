* `<<` 
all bit of binary of left   move to left and `符号位不变` 
left bit remove   and right bit add 0
`variate` == left * 2^number

* `>>` 
all bit of binary of left   move to right and `符号位不变` 
left bit add 0   and right bit remove
`variate` == left / 2^number

* `&` 
binary of left `与` binary of right
1 & 1 => 1
1 & 0 => 0
0 & 1 => 0
0 & 0 => 0

* `|` 
binary of left `或` binary of right
1 | 1 => 1
1 | 0 => 1
0 | 1 => 1
0 | 0 => 0

* `^` 
binary of left `异或` binary of right
1 ^ 0 => 1
0 ^ 1 => 1
0 ^ 0 => 0
1 ^ 1 => 0
