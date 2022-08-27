package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(UserID int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaings []Campaign
	
	err := r.db.Find(&campaings).Error
	if err != nil {
		return campaings, err
	}

	return campaings, nil
}

 func (r *repository) FindByUserID(UserID int) ([]Campaign, error) {
	var campaings []Campaign
	err := r.db.Where("user_id = ?", UserID).Find(&campaings).Error

	if err != nil {
		return campaings, err
	}
	
	return campaings, nil

 }
