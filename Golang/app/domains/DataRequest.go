package domains

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type DataRequest struct {
	StartDate string `form:"startDate"`
	EndDate   string `form:"endDate"`
	MinCount  int32  `form:"minCount"`
	MaxCount  int32  `form:"maxCount"`
}
type DataResponse struct {
	Id         int32     `json:"id"`
	CreatedAt  time.Time `json"created"`
	TotalMarks int       `json:"total"`
}
type JSONB []interface{}
type CasRecord struct {
	ID      int32     `json:"id" sql:"id"`
	Name    string    `json:"name" sql:"name"`
	Marks   JSONB     `json:"marks" sql:"marks"`
	Created time.Time `json:"created" sql:"created"`
}

func (CasRecord) TableName() string {
	return "cas_records"
}

// Value Marshal
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
