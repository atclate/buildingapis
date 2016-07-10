//************************************************************************//
// API "GoWorkshop": Application Resource Href Factories
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/bketelsen/buildingapis/exercises/21-goa/solution/design
// --out=$(GOPATH)/src/github.com/bketelsen/buildingapis/exercises/22-goa-logging-metrics/solution
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// CourseHref returns the resource href.
func CourseHref(id interface{}) string {
	return fmt.Sprintf("/api/courses/%v", id)
}

// RegistrationHref returns the resource href.
func RegistrationHref(id interface{}) string {
	return fmt.Sprintf("/api/registrations/%v", id)
}
