package model

type Gender string
type State string
type VehicleType string

// constants for Owner model
const (
	MALE   Gender = "male"
	FEMALE Gender = "female"
)

// constants for Vehicle model
const (
	ELECTRO  VehicleType = "electro"
	HYBRID   VehicleType = "hybrid"
	GASOLINE VehicleType = "gasoline"
	DIESEL   VehicleType = "diesel"
)

// constants for all models
const (
	ACTIVE   State = "active"
	INACTIVE State = "inactive"
)
