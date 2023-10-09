package usecase

import (
	"reflect"
	"testing"
)

func Test_FindAll(t *testing.T) {
	startDate := "2022 - 01 - 01"
	endDate := "2024 - 01 - 01"
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test Find All",
			want: "cGFzc3dvcmRpdHVSYWhhc2lhRGFuVGlkYWtCb2xlaGRpa2FzaXRhdQ==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got , err:= FindAll(startDate, endDate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("test decode = %v, want %v", got, tt.want)
			}
		})
	}
}
