package repository

import (
	"database/sql"

	"github.com/adrenalinchik/gowebapp/model"
	"github.com/adrenalinchik/gowebapp/log"
)

type VehicleRepo struct {
}

func InitVehicleRepo() (r *VehicleRepo) {
	return &VehicleRepo{}
}

func (r *VehicleRepo) Get(id int64) (vehicle *model.Vehicle, err error) {
	row := db.QueryRow("SELECT * FROM VEHICLE WHERE ID = ?", id)
	vehicle = new(model.Vehicle)
	err = row.Scan(&vehicle.Id, &vehicle.OwnerId, &vehicle.Model, &vehicle.Number, &vehicle.Type, &vehicle.State)
	if err != nil {
		log.Error.Println(err)
	}
	return
}

func (r *VehicleRepo) GetByOwner(id int64) (vehicles []*model.Vehicle, err error) {
	rows, err := db.Query("SELECT * FROM VEHICLE WHERE OWNER_ID = ?", id)
	if err != nil {
		log.Error.Println(err)
	}
	defer rows.Close()
	vehicles = make([]*model.Vehicle, 0)
	for rows.Next() {
		veh := new(model.Vehicle)
		err = rows.Scan(&veh.Id, &veh.OwnerId, &veh.Model, &veh.Number, &veh.Type, &veh.State)
		if err != nil {
			log.Error.Println(err)
		}
		vehicles = append(vehicles, veh)
	}
	return
}

// Get vehicle by number guarantee to return unique vehicle
func (r *VehicleRepo) GetByNumber(num string) (veh *model.Vehicle, err error) {
	row := db.QueryRow("SELECT * FROM VEHICLE WHERE NUMBER = ?", num)
	veh = new(model.Vehicle)
	err = row.Scan(&veh.Id, &veh.OwnerId, &veh.Model, &veh.Number, &veh.Type, &veh.State)
	if err == sql.ErrNoRows {
		log.Warning.Println(err)
	}
	if err != nil && err != sql.ErrNoRows {
		log.Error.Println(err)
	}
	return
}

func (r *VehicleRepo) Insert(vehicle *model.Vehicle) (id int64, err error) {
	res, err := db.Exec("INSERT INTO VEHICLE(OWNER_ID, MODEL, NUMBER, TYPE, STATE) VALUES(?,?,?,?,?)",
		vehicle.OwnerId, vehicle.Model, vehicle.Number, vehicle.Type, vehicle.State)
	if err != nil {
		log.Error.Println(err)
	}
	id, err = res.LastInsertId()
	if err != nil {
		log.Error.Println(err)
	}
	return
}

func (r *VehicleRepo) Update(vehicle *model.Vehicle) (id int64, err error) {
	_, err = db.Exec("UPDATE VEHICLE SET OWNER_ID = ?, TYPE = ?, NUMBER = ?, MODEL = ?, STATE = ? WHERE ID=?",
		vehicle.OwnerId, vehicle.Model, vehicle.Number, vehicle.Type, vehicle.State, vehicle.Id)
	id = vehicle.Id
	if err != nil {
		log.Error.Println(err)
	}
	return
}