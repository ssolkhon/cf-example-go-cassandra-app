# READ ME

### Create Service
```
$ cf create-service cassandra default cf-example-cass
```

### Get Dependencies
```
$ go get github.com/gocql/gocql
$ go get github.com/tools/godep
```

### Package Dependencies For Cloud Foundry
```
$ godep save
```

### Deploy
```
$ cf push
```