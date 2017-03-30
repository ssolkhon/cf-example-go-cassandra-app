package cassandra

import (
	"cf-example-go-cassandra-app/cf"
	"github.com/gocql/gocql"
	"strings"
)

func GetSession(services *cf.Services) (*gocql.Session, error) {
	/*
	   Create gocql session for accessing Cassandra cluster
	   Return gocql session
	*/
	myHosts := strings.Split(services.Cassandra[0].Credentials.Hosts, ",")
	cluster := gocql.NewCluster(myHosts[0])
	cluster.Keyspace = services.Cassandra[0].Credentials.Keyspace
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: services.Cassandra[0].Credentials.Username,
		Password: services.Cassandra[0].Credentials.Password,
	}
	cluster.Consistency = gocql.LocalOne

	return cluster.CreateSession()
}

func CreateTable(session *gocql.Session, tableName string) error {
	/*
	   Create table if it does not already exist
	   Return err if any
	*/
	q := `CREATE TABLE IF NOT EXISTS ` + tableName + ` (
        id varchar PRIMARY KEY,
        value varchar
        )`

	return session.Query(q).Exec()
}

func tableExists(session *gocql.Session, tableName string) bool {
	/*
	   Check if table exists
	   Return bool true if table exists
	*/
	q := `SELECT columnfamily_name
    FROM system.schema_columnfamilies
    WHERE keyspace_name='local_example_cass' AND columnfamily_name='` + tableName + `'`

	if numRows := session.Query(q).Iter().NumRows(); numRows > 0 {
		return true
	} else {
		return false
	}
}
