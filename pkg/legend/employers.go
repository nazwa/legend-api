package legend

import (
	"fmt"
	"net/url"
)

type EmployerModel struct {
	Id   int64
	Name string
}

// GetEmployers returns list of active employers at given location
func (this *ApiClient) GetEmployers(locationGuid string) (target []EmployerModel, err error) {
	mockFile := fmt.Sprintf("./mocks/employers-%s.json", locationGuid)
	if this.useMocks {
		err = this.getMockData(mockFile, &target)
		return
	}

	params := url.Values{}

	params.Add("LocationGUID", locationGuid)
	err = this.GetApiRequest("/Joining/Employers", params, &target)

	if this.saveMocks && err == nil {
		err = this.saveMockData(mockFile, target)
	}

	return
}
