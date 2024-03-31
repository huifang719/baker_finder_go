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

type Review struct {
	ID      uuid.UUID `json:"review_id"`
	BakerID uuid.UUID `json:"baker_id"`
	Review  string	`json:"review"`
	Rating  string	`json:"rating"`
	UserID  uuid.UUID	`json:"user_id"`
	UserName string	`json:"user_name"`
}

type User struct {
	ID             uuid.UUID `json:"user_id"`
	UserName       string	`json:"user_name"`
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

func databaseReviewtoReview(dbReview database.Review) Review {
	return Review{
		ID:      dbReview.ID,
		BakerID: dbReview.BakerID,
		Review:  dbReview.Review,
		Rating:  dbReview.Rating,
		UserID:  dbReview.UserID,
	}
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:       dbUser.ID,
		UserName: dbUser.UserName,
	}
}