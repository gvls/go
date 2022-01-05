## 2 `tag` 
If add `omitempty` to end of `new name`(`<反引号>json:"<new name>,omitempty"<反引号>`, can't add ` ` to side of `,`), When do `序列化` , `property` will ignore `default value`
If want or nested `struct` be ignore in `序列化` , we need add `omitempty` and change it as `pointer` 

### 3  point custom name for `json` or `yaml` 
Add `<反引号>json:"<new name>" yaml:"<new name> bson:"<new name>"<反引号>`(`tag标签` can be get by `反射`) to end of `data type` 

