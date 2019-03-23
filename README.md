
# Go REST on Cassandra

## Instructions

1. Run cassandra
	`docker run -d --name cassandra -p 7199:7199 -p 7000:7000 -p 7001:7001 -p 9160:9160 -p 9042:9042 cassandra`

2. Create Cassandra keyspace and table
   1. Open CQL connection
		```
		docker inspect cassandra | grep IPAddress
		docker exec -it cassandra cqlsh <IP_ADDRESS>
		```
   2. Create Cassandra keyspace
		`CREATE KEYSPACE golang WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};`
   3. Create Cassandra table
		```
		USE golang;
		CREATE TABLE users ( id UUID, firstname text, lastname text, age int, email text, city text, PRIMARY KEY (id) );
		```

3. Prepare environment
	```
	cd $GOPATH/src/github.com
	git clone git@github.com:bygui86/go-rest-cassandra.git
	cd bygui86/go-rest-cassandra
	go get ./...
	```

---

## Links

* [Tutorial](https://getstream.io/blog/building-a-performant-api-using-go-and-cassandra/)
