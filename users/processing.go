package users

import (
	"net/http"
	"strconv"
)

// FormToUser -- fills a User struct with submitted form data
// params:
// . r - request reader to fetch form data or url params (unused here)
// returns:
// . User struct if successful
// . array of strings of errors if any occur during processing
func FormToUser(r *http.Request) (User, []string) {
	var user User
	var errStr, ageStr string
	var errs []string
	var err error

	user.FirstName, errStr = processFormField(r, "firstname")
	errs = appendError(errs, errStr)
	user.LastName, errStr = processFormField(r, "lastname")
	errs = appendError(errs, errStr)
	user.Email, errStr = processFormField(r, "email")
	errs = appendError(errs, errStr)
	user.City, errStr = processFormField(r, "city")
	errs = appendError(errs, errStr)

	ageStr, errStr = processFormField(r, "age")
	if len(errStr) != 0 {
		errs = append(errs, errStr)
	} else {
		user.Age, err = strconv.Atoi(ageStr)
		if err != nil {
			errs = append(errs, "Parameter 'age' not an integer")
		}
	}
	return user, errs
}

// appendError -- allows us to append a string onto an existing array of strings
func appendError(errs []string, errStr string) []string {
	if len(errStr) > 0 {
		errs = append(errs, errStr)
	}
	return errs
}

// processFormField -- will take our request interface, and a string, and return the field data or an error in one of two return positions
func processFormField(r *http.Request, field string) (string, string) {
	// gorilla/mux has a PostFormValue() method defined for us to fetch form data from an HTTP request
	fieldData := r.PostFormValue(field)
	if len(fieldData) == 0 {
		return "", "Missing '" + field + "' parameter, cannot continue"
	}
	return fieldData, ""
}
