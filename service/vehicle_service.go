package service

import (
	"errors"

	"github.com/adrenalinchik/gowebapp/repository"
	"github.com/adrenalinchik/gowebapp/model"
	"github.com/adrenalinchik/gowebapp/log"
)

var vehicleRepo = repository.InitVehicleRepo()

type VehicleService struct {
}

func InitVehicleService() (s *VehicleService) {
	return &VehicleService{}
}

func (s *VehicleService) Create(v *model.Vehicle) (vehicle *model.Vehicle, err error) {
	v.Id, err = vehicleRepo.Insert(v)
	if err == nil {
		log.Info.Printf("Vehicle %d is saved to the db.", v.Id)
	}
	vehicle = v
	return
}

func (s *VehicleService) GetByOwner(id int64) (vehicles [] *model.Vehicle, err error) {
	vehicles, err = vehicleRepo.GetByOwner(id)
	if err != nil {
		log.Error.Println(err)
	} else {
		log.Info.Println("get all owner vehicles in the system")
	}
	return
}

func (s *VehicleService) GetByNumber(number string) (vehicle *model.Vehicle, err error) {
	vehicle, err = vehicleRepo.GetByNumber(number)
	if err != nil {
		log.Error.Println(err)
	} else {
		log.Info.Println("get vehicle by its number")
	}
	return
}

func (s *VehicleService) Update(v *model.Vehicle) (vehicle *model.Vehicle, err error) {
	if !v.Valid() {
		err = errors.New("vehicle validation is failed")
		vehicle = v
		return
	}
	if veh, _ := s.GetByNumber(v.Number); veh.Number == v.Number && veh.Id != v.Id {
		err = errors.New("vehicle with such number is already in the system")
		return
	}
	_, err = vehicleRepo.Update(v)
	if err == nil {
		log.Info.Printf("Vehicle %d is updated", v.Id)
	}
	return
}
