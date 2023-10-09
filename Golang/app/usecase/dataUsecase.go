package usecase

import (
	"encoding/json"
	"errors"
	"record/app/domains"
	"record/app/repocontainer"
)

type DataUsecase struct {
	repocontainer repocontainer.RepoContainer
}

func (l *DataUsecase) FindAll(start string, end string) (*[]domains.CasRecord, error) {
	return l.repocontainer.DataRepository.FindAll(start, end)
}

func (l *DataUsecase) InsertData(payload domains.CasRecord) (*domains.CasRecord, error) {
	return l.repocontainer.DataRepository.InsertData(payload)
}

func (l *DataUsecase) FilterData(payload *[]domains.CasRecord, maxCount int32, minCount int32) ([]domains.DataResponse, error) {
	var response []domains.DataResponse

	for _, y := range *payload {
		var marks []interface{}
		data, err := json.Marshal(y.Marks)
		if err != nil {
			return nil, errors.New("failed marshal")
		}
		err = json.Unmarshal(data, &marks)
		if err != nil {
			return nil, errors.New("failed marshal unmarshal")
		}
		totalMarks, err := l.FilterMarks(marks, maxCount, minCount)
		if err != nil {
			return nil, err
		}
		var res domains.DataResponse
		res.Id = y.ID
		res.CreatedAt = y.Created
		res.TotalMarks = totalMarks
		response = append(response, res)
	}
	return response, nil
}

func (i DataUsecase) FilterMarks(marks []interface{}, max int32, min int32) (int, error) {
	total := 0
	for _, val := range marks {
		if v, ok := val.(float64); ok {
			if v <= float64(max) && v >= float64(min) {
				total += 1
			}
		} else {
			return 0, errors.New("Error: Unable to convert value to integer")
		}
	}
	return total, nil
}

func NewDataUsecase(r *repocontainer.RepoContainer) *DataUsecase {
	return &DataUsecase{
		repocontainer: *r,
	}
}
