package business

import (
	"errors"
	"uniconnect/engine"
	"uniconnect/graph/model"
	"uniconnect/models"
	"uniconnect/utils"
)

func CreateNewBusiness(input model.CreateBusinessInput) (bool, error) {
	user, err := engine.FetchUserByAuthToken(input.Token)
	if err != nil {
		return false, err
	}

	newBusiness := models.Business{
		UserID:        user.ID,
		OwnerUserName: user.Name,
		Name:          input.Name,
		Website: func(site *string) string {
			if site == nil {
				return ""
			} else {
				return *input.Website
			}
		}(input.Website),
		Contact:     input.Contact,
		Description: input.Description,
		Location:    input.Location,
		Type:        input.Type,
		Image:       input.Image,
	}

	err = utils.DB.Create(&newBusiness).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func EditBusiness(input model.CreateBusinessInput, bizId string) (bool, error) {
	user, err := engine.FetchUserByAuthToken(input.Token)
	if err != nil {
		return false, err
	}

	business, err := engine.FetchBusinessById(bizId)
	if err != nil {
		return false, err
	}

	if business.UserID != user.ID {
		return false, errors.New("this business is not yours")
	}

	newBusiness := models.Business{
		Name: input.Name,
		Website: func(site *string) string {
			if site == nil {
				return ""
			} else {
				return *input.Website
			}
		}(input.Website),
		Contact:     input.Contact,
		Description: input.Description,
		Location:    input.Location,
		Type:        input.Type,
		Image:       input.Image,
	}
	err = utils.DB.Model(&business).Updates(newBusiness).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func FetchBusinessData(token, id string) (*model.Business, error) {
	_, err := engine.FetchUserByAuthToken(token)
	if err != nil {
		return nil, err
	}
	business, err := engine.FetchBusinessById(id)
	if err != nil {
		return nil, err
	}
	return business.CreateToGraphData(), nil
}

func FetchAllBusinesses(input model.FetchBusinessListInput) ([]*model.Business, error) {
	user, err := engine.FetchUserByAuthToken(input.Token)
	if err != nil {
		return nil, err
	}

	var businesses []models.Business

	if input.Mine != nil && *input.Mine {
		businesses, err = engine.FetchUserBusiness(user.ID)
		if err != nil {
			return nil, err
		}
		goto returnData
	}

	if input.SearchTerm != nil {
		businesses, err = engine.FetchBusinessesBySearchName(*input.SearchTerm)
		if err != nil {
			return nil, err
		}
		goto returnData
	}

	if input.Type != nil {
		if input.Type.IsValid() {
			businesses, err = engine.FetchBusinesses(&models.Business{Type: *input.Type}, 20)
			if err != nil {
				return nil, err
			}
			goto returnData
		} else {
			return nil, errors.New("invalid business type")
		}
	}

	err = utils.DB.Find(&businesses).Error
	if err != nil {
		return nil, err
	}
	goto returnData

returnData:
	gqlBusiness := make([]*model.Business, len(businesses))
	for i, biz := range businesses {
		gqlBusiness[i] = biz.CreateToGraphData()
	}
	return gqlBusiness, nil
}

func DeleteMyBusiness(token, bizId string) (bool, error) {
	user, err := engine.FetchUserByAuthToken(token)
	if err != nil {
		return false, err
	}
	business, err := engine.FetchBusinessById(bizId)
	if err != nil {
		return false, err
	}
	if business.UserID != user.ID {
		return false, errors.New("this is not your business")
	}
	err = utils.DB.Delete(&business).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
