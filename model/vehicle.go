package model

type Vehicle struct {
	Id      int64       `json:"id"`
	OwnerId int64       `json:"owner_id"`
	Model   string      `json:"model"`
	Number  string      `json:"number"`
	Type    VehicleType `json:"type"`
	State   State       `json:"state"`
}

func (vehicle *Vehicle) Valid() bool {
	return len(vehicle.Model) > 0 &&
		len(vehicle.Number) > 0 &&
		(vehicle.State == ACTIVE || vehicle.State == INACTIVE)
}
