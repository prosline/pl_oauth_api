package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/prosline/pl_logger/logger"
)

var (
	session *gocql.Session
)

func init() {
	// Connect to Cassandra cluster:
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		logger.Error("Connection to Cassandra failed", err)
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
