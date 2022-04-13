##  process
1. Execute all `case` (from up to down and from left to right)
2. **random** choose one `case` which execute success   to enter

3. If all case execute failure, execute content of `default`
4. If all case execute failure and haven't `default`, procedure will **阻塞wait** 

5. After out of `case` or `default`, exit `select`  

* random
time spent of `case` will effect randomness of choice
