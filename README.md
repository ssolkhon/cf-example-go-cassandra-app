# READ ME

### Create Service
```
$ cf create-service cassandra default cf-keystore
```

### Deploy
```
$ cf push
```

### Run Local
```
$ go run main.go
```

### Run Tests
```
$ cqlsh -u <username> -p <password> -f db/cassandra/test_data/data.cql
$ go test -v ./...
```