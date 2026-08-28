package user

import (
	"context"
	"errors"

	"github.com/nurgal1ev/booking-service/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Create(ctx context.Context, user *models.User) error {
	query := u.db.WithContext(ctx).Create(user).Error
	return query
}

func (u *UserRepo) FindByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := u.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []models.User{}, nil
		}
		return nil, err
	}
	return users, nil
}

func (u *UserRepo) Update(ctx context.Context, id int, firstname, lastname, username, email, password string) (*models.User, error) {
	err := u.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Updates(models.User{
		FirstName: firstname,
		LastName:  lastname,
		Username:  username,
		Email:     email,
		Password:  password,
	}).Error

	if err != nil {
		return nil, err
	}

	var user models.User
	err = u.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) Delete(ctx context.Context, id uint) error {
	result := u.db.WithContext(ctx).Delete(&models.User{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("property not found")
	}

	return nil
}
