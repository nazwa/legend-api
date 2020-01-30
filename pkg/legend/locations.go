package legend

type LocationHierarchy struct {
	Address1     string
	Address2     string
	Address3     string
	Address4     string
	City         string
	Country      string
	FacilityId   int64
	FriendlyName string
	Guid         string
	Latitude     float64
	LocationId   int64
	Longitude    float64
	LongName     string
	Name         string
	ParentGuid   string
	Postcode     string
	Region       string
	ShortName    string
	Type         int64
}

// GetAllLocations returns all available locations
func (this *ApiClient) GetAllLocations() (target []LocationHierarchy, err error) {
	if this.useMocks {
		err = this.getMockData("./mocks/locations.json", &target)
		return
	}
	err = this.GetApiRequest("/Locations/ALL", nil, &target)

	if this.saveMocks && err == nil {
		err = this.saveMockData("./mocks/locations.json", target)
	}

	return
}
