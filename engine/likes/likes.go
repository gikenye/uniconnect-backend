package likes

import (
	"uniconnect/engine"
	"uniconnect/graph/model"
	"uniconnect/models"
	"uniconnect/utils"

	uuid "github.com/satori/go.uuid"
)

func AddBusinessLike(token, bizId string) (bool, error) {
	user, err := engine.FetchUserByAuthToken(token)
	if err != nil {
		return false, err
	}
	biz, err := engine.FetchBusinessById(bizId)
	if err != nil {
		return false, err
	}

	err = utils.DB.Model(&models.Likes{}).Where(&models.Likes{UserID: user.ID, BusinessID: uuid.FromStringOrNil(bizId)}).First(nil).Error
	if err == nil {
		err = utils.DB.Where(&models.Likes{UserID: user.ID, BusinessID: biz.ID}).Delete(&models.Likes{}).Error
		if err != nil {
			return false, err
		}
		err = utils.DB.Model(&models.Business{}).Where("id = ?", biz.ID).Update("like_count", biz.LikeCount-1).Error
		if err != nil {
			return false, err
		}
		return false, nil
	}
	newLike := models.Likes{
		UserID:     user.ID,
		BusinessID: biz.ID,
	}
	err = utils.DB.Create(&newLike).Error
	if err != nil {
		return false, err
	}
	err = utils.DB.Model(&models.Business{}).Where("id = ?", biz.ID).Update("like_count", biz.LikeCount+1).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func FetchAllLikedBusinesses(token string) ([]*model.Business, error) {
	user, err := engine.FetchUserByAuthToken(token)
	if err != nil {
		return nil, err
	}

	businesses, err := engine.FetchLikes(user.ID)

	// var businesses []models.Business
	// err = utils.DB.Model(&models.Business{}).Joins("INNER JOIN likes ON likes.user_id = ? AND business.id = likes.business_id", user.ID).Scan(&businesses).Error
	if err != nil {
		return nil, err
	}

	gqlBusiness := make([]*model.Business, len(businesses))
	for i, biz := range businesses {
		gqlBusiness[i] = biz.CreateToGraphData()
	}
	return gqlBusiness, nil
}

func CheckBusinessLiked(token, bizId string) (bool, error) {
	user, err := engine.FetchUserByAuthToken(token)
	if err != nil {
		return false, err
	}
	err = utils.DB.Model(&models.Likes{}).Where(&models.Likes{UserID: user.ID, BusinessID: uuid.FromStringOrNil(bizId)}).First(nil).Error
	if err != nil {
		return false, nil
	}
	return true, nil
}
