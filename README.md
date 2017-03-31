# READ ME

### Create Service
```
$ cf create-service cassandra default cf-example-cass
```

### Get Dependencies
```
$ go get github.com/gocql/gocql
$ go get github.com/tools/godep
$ go get github.com/onsi/ginkgo/ginkgo
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
$ cql -u <username> -p <password> -c db/cassandra/test_data/data.cql
$ ginkgo -r
```