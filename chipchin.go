package chipchin

import (
	"log/slog"
)

func Main() {
	slog.Debug("chipchin", "test", true)

	person := Person{FirstName: "John", LastName: "Doe"}

	fullName := person.FullName()
	slog.Debug("full name", "fullName", fullName)

	person.ChangeLastName("Smith")
	slog.Debug("udated last name", "lastName", person.LastName)
}
