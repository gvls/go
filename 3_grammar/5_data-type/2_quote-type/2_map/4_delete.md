##  delete
###   delete **specific** key-value 
```go
delete(mapName, key)	// if key not exist, it not throw error
```

###   delete **all** key-value 
visit all and delete
or `mapName = make ...` and old `map` will be `gc` recycle

