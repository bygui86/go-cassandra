package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bygui86/go-cassandra/cassandra"
	"github.com/gocql/gocql"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var errs []string
	var gocqlUuid gocql.UUID

	user, errs := FormToUser(r)

	fmt.Println("Insert new user")

	// have we created a user correctly
	var created bool = false

	// if we had no errors from FormToUser, we will attempt to save our data to Cassandra
	if len(errs) == 0 {
		// generate a unique UUID for this user
		gocqlUuid = gocql.TimeUUID()

		// write data to Cassandra
		if err := cassandra.Session.Query(
			`INSERT INTO users (id, firstname, lastname, email, city, age) VALUES (?, ?, ?, ?, ?, ?)`,
			gocqlUuid, user.FirstName, user.LastName, user.Email, user.City, user.Age).Exec(); err != nil {
			errs = append(errs, err.Error())
		} else {
			created = true
		}
	}

	// depending on whether we created the user
	if created {
		// return the resource ID in a JSON payload
		fmt.Println("Created new User: user_id", gocqlUuid)
		json.NewEncoder(w).Encode(NewUserResponse{ID: gocqlUuid})
	} else {
		// return our errors
		fmt.Println("Failed creation of new user, errors", errs)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
