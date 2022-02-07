##  delete element
Like add element

###   append 
return slice **is** old slice
return array **is** old array
return len   **is** old len-1
return cap   **is** old cap
```go
append(slice1[:index], slice1[index+1:]...) 
```
