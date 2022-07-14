## 创建Go Module

1.删除本地go.mod

````
go mod init github.com/AuroraYolo/moody.git/
````


2.推送至代码仓库

3.增加新版本时,在仓库打新tag

## 想用本地文件怎么办

1.go.mod文件追加:
```
replace github.com/jeffail/tunny => xxx/xxx
```
2.go vendor 缓存到本地
```
go mod vendor 
go build -mod vendor
```

