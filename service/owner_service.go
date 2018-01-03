package service

import (
	"errors"
	"database/sql"
	"strconv"
	"fmt"

	"github.com/adrenalinchik/gowebapp/repository"
	"github.com/adrenalinchik/gowebapp/model"
	"github.com/adrenalinchik/gowebapp/log"
)

var ownerRepo = repository.InitOwnerRepo()

type OwnerService struct {
}

func InitOwnerService() (s *OwnerService) {
	return &OwnerService{}
}

func (s *OwnerService) Get(id int64) (owner *model.Owner, err error) {
	owner, err = ownerRepo.Get(id)
	strid := strconv.FormatInt(id, 10)
	if err == sql.ErrNoRows {
		err = errors.New("no owner found with id = " + strid)
		log.Error.Println(err)
		return
	}
	owner.Vehicles, err = vehicleRepo.GetByOwner(id)
	if err != nil {
		log.Error.Println(err)
	} else {
		log.Info.Println("get owner with id = " + strid)
	}
	return
}

func (s *OwnerService) GetByEmail(email string) (owner *model.Owner, err error) {
	owner, err = ownerRepo.GetByEmail(email)
	if err == sql.ErrNoRows {
		err = errors.New("no owner found with email = " + email)
		log.Error.Println(err)
		return
	}
	owner.Vehicles, err = vehicleRepo.GetByOwner(owner.Id)
	if err != nil {
		log.Error.Println(err)
	} else {
		log.Info.Println("get owner with email = " + email)
	}
	return
}

func (s *OwnerService) GetByState(state model.State) (owners [] *model.Owner, err error) {
	owners, err = ownerRepo.GetByState(state)
	if err != nil {
		log.Error.Println(err)
	} else {
		log.Info.Printf("get all %s owners in the system", string(state))
	}
	for i := range owners {
		o := owners[i]
		o.Vehicles, err = vehicleRepo.GetByOwner(o.Id)
		if err != nil {
			log.Error.Println(err)
		}
	}
	return
}

func (s *OwnerService) GetAll() (owners [] *model.Owner, err error) {
	owners, err = ownerRepo.GetAll()
	if err != nil {
		log.Error.Println(err)
	} else {
		log.Info.Println("get all owners in the system")
	}
	for i := range owners {
		o := owners[i]
		o.Vehicles, err = vehicleRepo.GetByOwner(o.Id)
		if err != nil {
			log.Error.Println(err)
		}
	}
	return
}

func (s *OwnerService) Create(owner *model.Owner) (o *model.Owner, err error) {
	if !owner.Valid() {
		err = errors.New("owner validation is failed")
		o = owner
		return
	}
	if own, _ := s.GetByEmail(owner.Email); own.Email == owner.Email {
		err = errors.New("owner with such email is already in the system")
		return
	}
	owner.Id, err = ownerRepo.Insert(owner)
	if err == nil {
		log.Info.Printf("owner %d is saved to the db.", owner.Id)
	}
	for i := range owner.Vehicles {
		vehicle := owner.Vehicles[i]
		v, err := vehicleRepo.GetByNumber(vehicle.Number)
		if err == sql.ErrNoRows {
			vehicle.OwnerId = owner.Id
			vehicle.Id, err = vehicleRepo.Insert(vehicle)
		} else {
			log.Warning.Printf("vehicle %d already exist", v.Id)
		}
	}
	o, err = s.Get(owner.Id)
	return
}

func (s *OwnerService) Update(owner *model.Owner) (o *model.Owner, err error) {
	if !owner.Valid() {
		err = errors.New("owner validation is failed")
		o = owner
		return
	}
	if own, _ := s.GetByEmail(owner.Email); own.Email == owner.Email && own.Id != owner.Id {
		err = errors.New("owner with such email is already in the system")
		return
	}
	owner.Id, err = ownerRepo.Update(owner)
	if err == nil {
		log.Info.Printf("owner %d is updated", owner.Id)
	}
	o = owner
	return
}

func (s *OwnerService) Delete(id int64) (err error) {
	owner, err := ownerRepo.Get(id)
	err = ownerRepo.Delete(id)
	if err == sql.ErrNoRows {
		err = fmt.Errorf("no owner found with id = %d", id)
		log.Error.Println(err)
		return
	}
	for i := range owner.Vehicles {
		veh := owner.Vehicles[i]
		veh.OwnerId = 0
		veh.State = model.INACTIVE
		vehicleRepo.Update(veh)
	}
	if err != nil {
		log.Error.Println(err)
	} else {
		log.Info.Printf("owner %d is deleted", id)
	}
	return
}
