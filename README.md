# READ ME

### Create Service
```
$ cf create-service cassandra default cf-example-cass
```

### Get Dependencies
```
$ go get github.com/gocql/gocql
```

### Package Dependencies For Cloud Foundry
```
$ godep save
```

### Deploy
```
$ cf push
```

### Run Tests
```
$ cqlsh -u <username> -p <password> -f db/cassandra/test_data/data.cql
$ go test -v ./...
```