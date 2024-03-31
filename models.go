package main

import (
	"time"

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
	Speciality string	`json:"speciality"`
	Creator   uuid.UUID `json:"creator"`
}

type Review struct {
	ID      uuid.UUID `json:"review_id"`
	BakerID uuid.UUID `json:"baker_id"`
	Review  string	`json:"review"`
	Rating  string	`json:"rating"`
	UserID  uuid.UUID	`json:"user_id"`
	UserName string	`json:"user_name"`
	CreatedAt time.Time	`json:"created_at"`
	BakerName string	`json:"baker_name"`
}

type User struct {
	ID             uuid.UUID `json:"user_id"`
	UserName       string	`json:"user_name"`
	UserType       string	`json:"user_type"`
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
		Speciality: dbBaker.Specialty,
		Creator:   dbBaker.Creator,
	}
}	  

func databaseReviewtoReview(dbReview database.Review, dbBaker database.Baker, dbUser database.User ) Review {
	return Review{
		ID:      dbReview.ID,
		BakerID: dbReview.BakerID,
		Review:  dbReview.Review,
		Rating:  dbReview.Rating,
		UserID:  dbReview.UserID,
		CreatedAt: dbReview.CreatedAt,
		BakerName: dbBaker.Name,
		UserName: dbUser.UserName,
	}
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:       dbUser.ID,
		UserName: dbUser.UserName,
		UserType: dbUser.UserType,
	}
}