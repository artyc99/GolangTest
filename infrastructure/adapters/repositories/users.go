package repositories

import (
	"EchoAPI/core/models"
	"EchoAPI/core/repositories"
	"context"
)

type users struct {
	postgresDB GormDB
}

func Users(db GormDB) repositories.UsersI {
	return &users{
		postgresDB: db,
	}
}

func (r *users) Create(ctx context.Context, user models.User) (err error) {
	db := r.postgresDB.WithContext(ctx)

	err = db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *users) GetList(ctx context.Context, props repositories.FilterProps) (users []models.User, total int64, err error) {
	db := r.postgresDB.WithContext(ctx)

	reqTx := db.Model(&models.User{})

	if props.FromAge != nil && props.ToAge != nil {
		reqTx = reqTx.Where("age BETWEEN ? AND ?", props.FromAge, props.ToAge)
	} else if props.FromAge != nil {
		reqTx = reqTx.Where("age >= ?", props.FromAge)
	} else if props.ToAge != nil {
		reqTx = reqTx.Where("age <= ?", props.ToAge)
	}

	if props.FromDate != nil && props.ToDate != nil {
		reqTx = reqTx.Where("recording_date BETWEEN ? AND ?", props.FromDate, props.ToDate)
	} else if props.FromDate != nil {
		reqTx = reqTx.Where("recording_date >= ?", props.FromDate)
	} else if props.ToDate != nil {
		reqTx = reqTx.Where("recording_date <= ?", props.ToDate)
	}

	err = reqTx.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = reqTx.Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
