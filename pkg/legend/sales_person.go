package legend

import (
	"net/url"
	"strconv"
)

type SalesPerson struct {
	Id        int64
	FirstName string
	LastName  string
}

// GetSalesPersonDetails returns a list of sales-people at location
func (this *ApiClient) GetSalesPersonDetails(locationId int64) (target []SalesPerson, err error) {
	if this.useMocks {
		err = this.getMockData("./mocks/sales-person-"+strconv.FormatInt(locationId, 10)+".json", &target)
		return
	}

	params := url.Values{}

	params.Add("locationid", strconv.FormatInt(locationId, 10))
	err = this.GetApiRequest("/Joining/SalesPersonDetails", params, &target)

	if this.saveMocks && err == nil {
		err = this.saveMockData("./mocks/sales-person-"+strconv.FormatInt(locationId, 10)+".json", target)
	}

	return
}
