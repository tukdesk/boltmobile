#### BoltMobile
##### 目的

对 [boltdb](https://github.com/boltdb/bolt) 进行封装， 使之能够通过 [gomobile](https://github.com/golang/mobile) 编译出供移动平台使用的库

##### 使用方法

```
go get github.com/tukdesk/boltmobile
gomobile bind -target ios github.com/tukdesk/boltmobile
```

##### 说明

未来可能考虑对以下方法进行封装：  

```
bolt.DB.Stats

bolt.Tx.Stats
bolt.Tx.Writeto
bolt.Tx.CopyFile

bolt.Bucket.Root
bolt.Bucket.Stats
```

目前暂不考虑进行封装的方法：

```
bolt.DB.Sync
bolt.DB.Info

bolt.Tx.Check
bolt.Tx.Copy
bolt.Tx.Page
```

##### TODO
-   test cases