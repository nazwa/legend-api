package legend

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type MembershipPriceDetailsModel struct {
	CycleFee                   decimal.Decimal
	CycleFeeProcessType        int64
	EnrollmentFee              decimal.Decimal
	EnrollmentFeeProcessType   int64
	FirstBillDate              LegendTime
	HasEligibilities           bool
	InductionFee               decimal.Decimal
	InductionFeeProcessType    int64
	MonthInHandFee             decimal.Decimal
	MonthInHandFeeProcessType  int64
	MonthInHandProcessType     int64
	ProrateFee                 decimal.Decimal
	ProrateFeeProcessType      int64
	UpFrontCycleFee            decimal.Decimal
	UpfrontCycleFeeProcessType int64
	WaiveProrateFee            bool
	Discounts                  []DiscountPriceDetails
	FutureDiscounts            []FutureDiscountPriceDetails
}

// GetPriceDetails fetches base price for a given membership price at a location
func (this *ApiClient) GetPriceDetails(locationId int64, priceTypeId int64, startDate time.Time, promotionCode string) (target MembershipPriceDetailsModel, err error) {
	params := url.Values{}

	params.Add("PriceTypeId", strconv.FormatInt(priceTypeId, 10))
	params.Add("StartDate", startDate.Format(LEGEND_DATE_FORMAT))

	if promotionCode != "" {
		params.Add("PromotionCode", promotionCode)
	}

	err = this.GetApiRequest(fmt.Sprintf("/Joining/MembershipTypes/%d/PriceDetails", locationId), params, &target)

	return
}
