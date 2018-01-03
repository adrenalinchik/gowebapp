package repository

import (
	"github.com/adrenalinchik/gowebapp/model"
	"github.com/adrenalinchik/gowebapp/log"
)

type OwnerRepo struct {
}

func InitOwnerRepo() (r *OwnerRepo) {
	return &OwnerRepo{}
}

func (r *OwnerRepo) Get(id int64) (owner *model.Owner, err error) {
	row := db.QueryRow("SELECT * FROM OWNER WHERE ID = ?", id)
	owner = new(model.Owner)
	err = row.Scan(&owner.Id, &owner.Firstname, &owner.Lastname, &owner.Gender, &owner.Dob, &owner.Email, &owner.State)
	if err != nil {
		log.Error.Println(err)
	}
	return
}

func (r *OwnerRepo) GetByEmail(email string) (owner *model.Owner, err error) {
	row := db.QueryRow("SELECT * FROM OWNER WHERE EMAIL = ?", email)
	owner = new(model.Owner)
	err = row.Scan(&owner.Id, &owner.Firstname, &owner.Lastname, &owner.Gender, &owner.Dob, &owner.Email, &owner.State)
	if err != nil {
		log.Error.Println(err)
	}
	return
}

func (r *OwnerRepo) GetByState(state model.State) (owners [] *model.Owner, err error) {
	rows, err := db.Query("SELECT * FROM OWNER WHERE STATE = ?", state)
	if err != nil {
		log.Error.Println(err)
	}
	defer rows.Close()
	owners = make([]*model.Owner, 0)
	for rows.Next() {
		owner := new(model.Owner)
		err = rows.Scan(&owner.Id, &owner.Firstname, &owner.Lastname, &owner.Gender, &owner.Dob, &owner.Email, &owner.State)
		if err != nil {
			log.Error.Println(err)
		}
		owners = append(owners, owner)
	}
	return
}

func (r *OwnerRepo) GetAll() (owners [] *model.Owner, err error) {
	rows, err := db.Query("SELECT * FROM OWNER")
	if err != nil {
		log.Error.Println(err)
	}
	defer rows.Close()
	owners = make([]*model.Owner, 0)
	for rows.Next() {
		owner := new(model.Owner)
		err = rows.Scan(&owner.Id, &owner.Firstname, &owner.Lastname, &owner.Gender, &owner.Dob, &owner.Email, &owner.State)
		if err != nil {
			log.Error.Println(err)
		}
		owners = append(owners, owner)
	}
	return
}

func (r *OwnerRepo) Insert(owner *model.Owner) (id int64, err error) {
	res, err := db.Exec("INSERT INTO OWNER(FIRSTNAME, LASTNAME, GENDER, DOB, EMAIL, STATE) VALUES(?,?,?,?,?,?)",
		owner.Firstname, owner.Lastname, owner.Gender, owner.Dob, owner.Email, owner.State)
	if err != nil {
		log.Error.Println(err)
	}
	id, err = res.LastInsertId()
	if err != nil {
		log.Error.Println(err)
	}
	return
}

func (r *OwnerRepo) Update(owner *model.Owner) (id int64, err error) {
	_, err = db.Exec("UPDATE OWNER SET FIRSTNAME = ?, LASTNAME = ?, GENDER = ?, DOB = ?, EMAIL = ?, STATE = ? WHERE ID=?",
		owner.Firstname, owner.Lastname, owner.Gender, owner.Dob, owner.Email, owner.State, owner.Id)
	id = owner.Id
	if err != nil {
		log.Error.Println(err)
	}
	return
}

func (r *OwnerRepo) Delete(id int64) (err error) {
	_, err = db.Exec("DELETE FROM OWNER WHERE ID=?", id)
	if err != nil {
		log.Error.Println(err)
	}
	return
}