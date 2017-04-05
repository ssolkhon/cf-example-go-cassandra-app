package cassandra

import (
	"cf-keystore/cf"
	"errors"
	"github.com/gocql/gocql"
	"strings"
)

func GetSession(c cf.CassandraService) (*gocql.Session, error) {
	/*
	   Create gocql session for accessing Cassandra cluster
	   Return gocql session
	*/
	myHosts := strings.Split(c.Credentials.Hosts, ",")
	cluster := gocql.NewCluster(myHosts[0])
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: c.Credentials.Username,
		Password: c.Credentials.Password,
	}
	cluster.Consistency = gocql.LocalOne

	return cluster.CreateSession()
}

func CreateTable(session *gocql.Session, keyspace string, table string) error {
	/*
	   Create table if it does not already exist
	   Return err if any
	*/
	q := `CREATE TABLE IF NOT EXISTS "` + keyspace + `"."` + table + `" (
        id varchar PRIMARY KEY,
        value varchar
        )`

	return session.Query(q).Exec()
}

func CreateRow(session *gocql.Session, keyspace string, table string, key string, value string) error {
	/*
	   Create row in table or overwrite existing value
	   Return err if any
	*/
	q := `INSERT INTO "` + keyspace + `"."` + table + `" (id, value)
      VALUES ('` + key + `', '` + value + `')`

	if exists := tableExists(session, keyspace, table); exists == true {
		return session.Query(q).Exec()
	} else {
		return errors.New(table + " does not exist - please create it.")
	}
}

func GetRow(session *gocql.Session, keyspace string, table string, key string) (string, error) {
	/*
	   Get row in table
	   Return err if any
	*/
	var myValue string
	q := `SELECT value
      FROM "` + keyspace + `"."` + table + `"
      WHERE id='` + key + `'`

	if exists := tableExists(session, keyspace, table); exists == true {
		err := session.Query(q).Scan(&myValue)
		return myValue, err
	} else {
		err := errors.New(table + " does not exist - please create it.")
		return myValue, err
	}
}

func tableExists(session *gocql.Session, keyspace string, table string) bool {
	/*
	   Check if table exists
	   Return bool true if table exists
	*/
	q := `SELECT columnfamily_name
    FROM system.schema_columnfamilies
    WHERE keyspace_name='` + keyspace + `' AND columnfamily_name='` + table + `'`

	if numRows := session.Query(q).Iter().NumRows(); numRows > 0 {
		return true
	} else {
		return false
	}
}
