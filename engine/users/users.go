package users

import (
	"errors"
	"uniconnect/engine"
	"uniconnect/graph/model"
	"uniconnect/models"
	"uniconnect/utils"
)

func CheckUserNameAvaliability(username string) (bool, error) {
	_, err := engine.FetchUser(&models.User{UserName: username}, false)
	if err == nil {
		return false, errors.New("username already exists")
	}
	return true, nil
}
func ValidateRegistration(input model.RegisterInput) (*model.AuthPayload, error) {
	if !(input.Password == input.Confirmpassword) {
		return nil, errors.New("passwords do not match")
	}
	_, err := engine.FetchUser(&models.User{Email: input.Email}, false)
	if err == nil {
		return nil, errors.New("email already exists")
	}
	_, err = engine.FetchUser(&models.User{UserName: input.Username}, false)
	if err == nil {
		return nil, errors.New("username already exists")
	}

	password, err := utils.HashString(input.Password)
	if err != nil {
		return nil, errors.New("internal encryption error")
	}

	newUser := models.User{
		Name:     input.Fullname,
		Email:    input.Email,
		UserName: input.Username,
		Password: password,
	}

	err = utils.DB.Create(&newUser).Error
	if err != nil {
		return nil, err
	}
	return engine.GenerateToken(newUser.ID)
}

func Login(input model.LoginInput) (*model.AuthPayload, error) {
	var user *models.User
	var err error

	user, err = engine.FetchUser(&models.User{Email: input.EmailorUsername}, false)
	if err != nil {
		user, err = engine.FetchUser(&models.User{UserName: input.EmailorUsername}, false)
		if err != nil {
			return nil, err
		}
	}

	if !utils.CompareHashedString(user.Password, input.Password) {
		return nil, engine.ErrWrongCreds
	}
	return engine.GenerateToken(user.ID)
}

func FetchProfile(token string) (*model.User, error) {
	user, err := engine.FetchUserByAuthToken(token)
	if err != nil {
		return nil, err
	}
	return user.CreateToGraphData(), nil
}

func ChangePassword(input model.ChangePasswordInput) (bool, error) {
	if input.ConfirmNewPassword != input.NewPassword {
		return false, errors.New("confirm password does not match")
	}
	if input.OldPassword == input.NewPassword {
		return false, errors.New("please update to a different password")
	}
	user, err := engine.FetchUserByAuthToken(input.Token)
	if err != nil {
		return false, err
	}
	if !utils.CompareHashedString(user.Password, input.OldPassword) {
		return false, errors.New("wrong old password")
	}
	hashedNewPass, err := utils.HashString(input.NewPassword)
	if err != nil {
		return false, errors.New("internal encryption error")
	}

	err = utils.DB.Model(&user).Update("password", hashedNewPass).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func VerifyEmail(token, otp string) (bool, error) {
	user, err := engine.FetchUserByAuthToken(token)
	if err != nil {
		return false, err
	}
	err = utils.DB.Model(&user).Update("verified", true).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
