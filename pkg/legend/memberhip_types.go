package legend

import (
	"net/url"
	"strconv"
)

type FeeType int64

const (
	FeeTypeUnknown     FeeType = 0
	FeeTypeDirectDebit FeeType = 1
	FeeTypePayInFull   FeeType = 2
	FeeTypeCreditCard  FeeType = 3
)

func (f FeeType) String() string {
	switch f {
	case FeeTypeDirectDebit:
		return "Direct Debit"
	case FeeTypePayInFull:
		return "Pay in Full"
	case FeeTypeCreditCard:
		return "Credit Card"
	default:
		return "Unknown"
	}
}

type DiscountPriceDetails struct {
	Amount                    float64
	CycleDiscountActiveMonths int64
	Description               string
	ForProcessType            int64
	ProcessType               int64
	PromoCode                 string
	PromotionID               int64
	PromotionName             string
}

type AgreementChannelEventModel struct {
	EventCode string
	Prompt    string
}

type FutureDiscountPriceDetails struct {
	Amount                    float64
	CycleDiscountActiveMonths int64
	Description               string
	ForProcessType            int64
	ProcessType               int64
	PromoCode                 string
	PromotionID               int64
	PromotionName             string
}

type MembershipPriceTypeModel struct {
	Id                         int64
	ActiveFrom                 LegendTime
	ActiveTo                   LegendTime
	AgreementChannelID         int64
	CycleFee                   float64
	CycleFeeProcessType        int64
	Detail                     string
	DurationInMonths           int64
	EnrollmentFee              float64
	EnrollmentFeeProcessType   int64
	FeeType                    FeeType
	HasEligibilities           bool
	InductionFee               float64
	InductionFeeProcessType    int64
	MembershipDurationInDays   int64
	MonthInHandFee             float64
	MonthInHandFeeProcessType  int64
	MonthInHandProcessType     int64
	MonthsInHand               int64
	Name                       string
	ProrateFee                 float64
	ProrateFeeProcessType      int64
	UpFrontCycleFee            float64
	UpfrontCycleFeeProcessType int64
	WaiveProrateFee            bool
	WebDescription             string
	Discounts                  []DiscountPriceDetails
	Events                     []AgreementChannelEventModel
	FutureDiscounts            []FutureDiscountPriceDetails
}

type MembershipTypeModel struct {
	Id               int64
	Detail           string
	HasEligibilities bool
	Name             string
	WebDescription   string
	Prices           []MembershipPriceTypeModel
}

type MembershipTypeReturnModel struct {
	LocationId     int64
	LocationName   string
	MembershipType []MembershipTypeModel
}

// GetMembershipTypes reads available memberships at a given location
func (this *ApiClient) GetMembershipTypes(startDate LegendTime, endDate LegendTime, locationId int64, channelId int64) (target []MembershipTypeReturnModel, err error) {
	if this.useMocks {
		err = this.getMockData("./mocks/membership-types-"+strconv.FormatInt(locationId, 10)+".json", &target)
		return
	}

	params := url.Values{}

	params.Add("startDate", startDate.FormatShortDate())
	params.Add("endDate", endDate.FormatShortDate())

	if locationId != -1 {
		params.Add("locationId", strconv.FormatInt(locationId, 10))
	}
	if channelId != -1 {
		params.Add("channelId", strconv.FormatInt(channelId, 10))
	}

	err = this.GetApiRequest("/Joining/MembershipTypes", params, &target)

	if this.saveMocks && err == nil {
		err = this.saveMockData("./mocks/membership-types-"+strconv.FormatInt(locationId, 10)+".json", target)
	}

	return
}
