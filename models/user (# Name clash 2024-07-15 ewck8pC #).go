package models

import "uniconnect/graph/model"

type User struct {
	Base
	Name         string
	Email        string
	Password     string
	UserName     string
	ProfilePhoto string
	Verified     bool
}

func (u User) CreateToGraphData() *model.User {
	return &model.User{
		ID:           u.ID.String(),
		Name:         u.Name,
		Email:        u.Email,
		Username:     u.UserName,
		ProfilePhoto: u.ProfilePhoto,
		Verified:     u.Verified,
	}
}
