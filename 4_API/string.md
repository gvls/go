## 2 system function for string
`string` can't be change
`function` will return new `string` 

### 3  get length
`len(<type>) int` 

`string` : get length by byte 

### 3  select `string` by byte
`<string>[index]` 

### 3  `string` => `slice` 
`[]rune(<string>) slice` 

### 3  select `slice` by character
`<slice>[index]` 

### 3  `string` => `整数` 
`strconv.Atoi(<string>) (int, error)` 

`error` can be use in check `data type` of `variate` 

### 3  `整数` => `string` 
`strconv.Itoa(<int>) string` 

### 3  `string` => `[]byte` 
`[]byte(<string>) byte` 

### 3  `[]byte` => `string` 
`string(<[]byte>) string` 

### 3  if exist `sub string` in `string` 
`strings.Contains(<main string>, <sub string>) bool` 

### 3  count of `sub string` in `string` 
`strings.Count(<main string>, <sub string> int` 

### 3  `index` of first appear of `sub string` in `string` 
`strings.Index(<main string>, <sub string>) int` 
If can't find `sub string` in `string` , return -1

### 3  `index` of last appear of `sub string` in `string` 
`strings.LastIndex(<main string>, <sub string>) int` 
If can't find `sub string` in `string` , return -1

### 3  replace `sub string` in `string` 
`strings.Replace(<main string>, <sub string>, <new sub string>, <次数>) string` 
If `次数` == -1, it is replace all `sub string` 

### 3  if equal between `string` and `string` 
`<string> == <string> bool` 

### 3  if equal between `string` and `string` by not case sensitive
`strings.EqualFold(<string>, <string>) bool` 

### 3  divide `string` to `[]string` by character
`strings.Split(<string>, <character>) []string` 

### 3  upper all alphabet of `string` 
`strings.ToLower(<string>) string` 

### 3  lower all alphabet of `string` 
`strings.ToUpper(<string>) string` 

### 3  remove `space` in side of `string` 
`strings.TrimSpace(<string>) string` 

### 3  remove character in side of `string` 
`strings.Trim(<string>, <characters>) string` 
characters is consist of 1 ~ xxx character

### 3  remove character in left of `string` 
`strings.TrimLeft(<string>, <characters>) string` 
characters is consist of 1 ~ xxx character

### 3  remove character in right of `string` 
`strings.TrimRight(<string>, <characters>) string` 
characters is consist of 1 ~ xxx character

### 3  if `string` begin with `sub string` 
`strings.HasPrefix(<main string>, <sub string>) bool` 

### 3  if `string` end with `sub string` 
`strings.HasSuffix(<main string>, <sub string>) bool` 
