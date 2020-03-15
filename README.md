
# Go REST on Cassandra

## Instructions

1. Prepare environment
	```
	cd $GOPATH/src/github.com
	git clone git@github.com:bygui86/go-cassandra.git
	cd bygui86/go-cassandra
	go get ./...
	```

2. Run cassandra
	```
	docker run -d --name cassandra -p 7199:7199 -p 7000:7000 -p 7001:7001 -p 9160:9160 -p 9042:9042 cassandra
	```

3. Create Cassandra keyspace and table
   1. Open CQL connection
		```
		docker exec -it cassandra cqlsh $(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' cassandra)
		```
   2. Create Cassandra keyspace
		```
		CREATE KEYSPACE golang WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};
		```
   3. Create Cassandra table
		```
		USE golang;
		CREATE TABLE users ( id UUID, firstname text, lastname text, age int, email text, city text, PRIMARY KEY (id) );
		```

4. Run application
	```
	cd $GOPATH/src/github.com/bygui86/go-cassandra
	go run main.go
	```

5. Test some REST calls (see endpoints list below)

---

## REST endpoints

* `GET http://localhost:8080/` heart beat
* `POST http://localhost:8080/users/new` insert new user
* `GET http://localhost:8080/users` get all users
* `GET http://localhost:8080/users/{uuid}` get user by uuid

`PLEASE NOTE`: See the postman collection for samples of the user JSON structure

---

## Links

* [Tutorial](https://getstream.io/blog/building-a-performant-api-using-go-and-cassandra/)
