package model

import (
	"time"
	"strings"
)

type Owner struct {
	Id        int64      `json:"id"`
	Firstname string     `json:"firstname"`
	Lastname  string     `json:"lastname"`
	Gender    Gender     `json:"gender"`
	Dob       time.Time  `json:"dob"`
	Email     string     `json:"email"`
	State     State      `json:"state"`
	Vehicles  []*Vehicle `json:"vehicles"`
}

func (owner *Owner) Valid() bool {
	start, _ := time.Parse(time.RFC822, "01 Jan 1900 10:00 UTC")
	return len(owner.Firstname) > 0 &&
		len(owner.Lastname) > 0 &&
		(owner.Gender == MALE || owner.Gender == FEMALE) &&
		owner.Dob.After(start) &&
		len(owner.Email) > 3 &&
		strings.Contains(owner.Email, "@") &&
		strings.Contains(owner.Email, ".") &&
		(owner.State == ACTIVE || owner.State == INACTIVE)
}
