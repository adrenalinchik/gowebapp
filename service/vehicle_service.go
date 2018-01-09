package service

import (
	"github.com/adrenalinchik/gowebapp/repository"
	"github.com/adrenalinchik/gowebapp/model"
	"github.com/adrenalinchik/gowebapp/log"
	"fmt"
	"strconv"
	"database/sql"
)

var vehicleRepo = repository.InitVehicleRepo()

var ownerService = InitOwnerService()

type VehicleService struct {
}

func InitVehicleService() (s *VehicleService) {
	return &VehicleService{}
}

func (s *VehicleService) Get(id int64) (vehicle *model.Vehicle, err error) {
	vehicle, err = vehicleRepo.Get(id)
	strid := strconv.FormatInt(id, 10)
	if err == sql.ErrNoRows {
		err = fmt.Errorf("no vehicle found with id = " + strid)
		log.Error.Println(err)
		return
	}
	log.Info.Println("get owner with id = " + strid)
	return
}

func (s *VehicleService) GetByOwner(id int64) (vehicles [] *model.Vehicle, err error) {
	_, err = ownerService.Get(id)
	if err != nil {
		return
	}
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

func (s *VehicleService) Create(v *model.Vehicle) (vehicle *model.Vehicle, err error) {
	if !v.Valid() {
		err = fmt.Errorf("vehicle validation is failed")

		return
	}
	if veh, _ := s.GetByNumber(v.Number); veh.Number == v.Number && veh.Id != v.Id {
		err = fmt.Errorf("vehicle with such number is already in the system")
		return
	}
	v.Id, err = vehicleRepo.Insert(v)
	if err == nil {
		log.Info.Printf("vehicle %d is saved to the db", v.Id)
	}
	vehicle, err = s.Get(v.Id)
	return
}

func (s *VehicleService) Update(v *model.Vehicle) (vehicle *model.Vehicle, err error) {
	if !v.Valid() {
		err = fmt.Errorf("vehicle validation is failed")
		return
	}
	if veh, _ := s.GetByNumber(v.Number); veh.Number == v.Number && veh.Id != v.Id {
		err = fmt.Errorf("vehicle with such number is already in the system")
		return
	}
	_, err = vehicleRepo.Update(v)
	if err == nil {
		log.Info.Printf("vehicle %d is updated", v.Id)
	}
	vehicle, err = s.Get(v.Id)
	return
}
