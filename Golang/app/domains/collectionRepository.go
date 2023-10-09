package domains

type DataRepository interface {
	FindAll(startDate string, endDate string) (*[]CasRecord, error)
	InsertData(payload CasRecord) (*CasRecord, error)
}
