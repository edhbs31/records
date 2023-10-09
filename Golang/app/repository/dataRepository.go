package repository

import (
	"errors"
	"fmt"
	"record/app/domains"
	"time"

	"gorm.io/gorm"
)

type DataRepository struct {
	db *gorm.DB
}

func (i DataRepository) FindAll(startDateStr string, endDateStr string) (*[]domains.CasRecord, error) {
	var items *[]domains.CasRecord = nil
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		// handle error
	}
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, errors.New("start date error parse")
	}
	result := i.db.Debug().Where("created<=? AND created>=? ", endDate, startDate).Find(&items)
	if result == nil {
		return nil, errors.New("not found")
	}
	return items, nil
}

func (i DataRepository) InsertData(payload domains.CasRecord) (*domains.CasRecord, error) {
	payload.Created = time.Now()
	err := i.db.Create(&payload).Error
	if err != nil {
		fmt.Print("Tjis is error nya uasaa", err.Error())
		return nil, errors.New("failed insert")
	}
	return &payload, nil
}

func NewDataRepository(c *gorm.DB) domains.DataRepository {
	return &DataRepository{db: c}
}
