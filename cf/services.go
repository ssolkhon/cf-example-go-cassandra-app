package cf

type CassandraService []struct {
	Name        string
	Label       string
	Tags        []string
	Plan        string
	Credentials struct {
		Username           string
		Password           string
		Keyspace           string
		Port               int
		Hosts              string
		FirewallAllowRules string
	}
}

type Services struct {
	Cassandra CassandraService
}
