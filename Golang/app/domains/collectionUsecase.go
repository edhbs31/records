package domains

type DataUsecase interface {
	FindAll(startDate string, endDate string) (*[]CasRecord, error)
	FilterData(payload *[]CasRecord, maxCount int32, minCount int32) ([]DataResponse, error)
	InsertData(payload CasRecord) (*CasRecord, error)
}
