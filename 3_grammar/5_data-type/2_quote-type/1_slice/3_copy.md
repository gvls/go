## 2 copy
copy **element** of `array2` to `array1` 
```go
// slice1 must be initialize
copy(slice1, slice2)
```

* if `len1`   `小于`   `len2` 
~~`array1` may lose some element of `array2`~~

* if `len1`   `大于`   `len2` 
end element of `array1` keep old value


