package main

import (
	"github.com/google/uuid"
	"github.com/huifang719/baker_finder_go/internal/database"
)

type Baker struct {
	ID        uuid.UUID `json:"baker_id"`
	Img       string 	`json:"img"`
	Name      string	`json:"name"`
	Address   string 	`json:"address"`
	Suburb    string 	`json:"suburb"`
	Postcode  string 	`json:"postcode"`
	Contact   string 	`json:"contact"`
	Specialty string	`json:"specialty"`
	Creator   uuid.UUID `json:"creator"`
}

func databaseBakertoBaker(dbBaker database.Baker) Baker {
	return Baker{
		ID:        dbBaker.ID,
		Img:       dbBaker.Img,
		Name:      dbBaker.Name,
		Address:   dbBaker.Address,
		Suburb:    dbBaker.Suburb,
		Postcode:  dbBaker.Postcode,
		Contact:   dbBaker.Contact,
		Specialty: dbBaker.Specialty,
		Creator:   dbBaker.Creator,
	}
}	  