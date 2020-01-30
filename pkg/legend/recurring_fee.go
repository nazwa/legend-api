package legend

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type BillingMethod int64

const (
	BillingMethodUnknown     BillingMethod = 0
	BillingMethodDirectDebit BillingMethod = 1
	BillingMethodCreditCard  BillingMethod = 2
	BillingMethodPayInFull   BillingMethod = 8
)

type RecurringFeeModel struct {
	Id                 int64
	ActiveDate         LegendTime
	AdminFee           decimal.Decimal
	Autobill           bool
	BillingCycleId     int64
	BillingMethodId    BillingMethod
	ClubRestrictionId  int64
	CycleFee           decimal.Decimal
	Description        string
	GroupRestrictionId int64
	InactiveDate       LegendTime
	IsEnterprise       bool
	LocationId         int64
	Name               string
	UsageMonths        int64
	WaiveProrateFee    bool
}

// GetRecurringFee returns List of recurring fees (active only)
func (this *ApiClient) GetRecurringFee(facilityGuid string, billingMethodId BillingMethod) (target []RecurringFeeModel, err error) {
	mockFile := fmt.Sprintf("./mocks/recurring-fee-%s-%d.json", facilityGuid, billingMethodId)
	if this.useMocks {
		err = this.getMockData(mockFile, &target)
		return
	}

	params := url.Values{}

	params.Add("FacilityGUID", facilityGuid)
	params.Add("BillingMethodId", strconv.FormatInt(int64(billingMethodId), 10))
	err = this.GetApiRequest("/Joining/RecurringFee", params, &target)

	if this.saveMocks && err == nil {
		err = this.saveMockData(mockFile, target)
	}

	return
}

func (bm BillingMethod) String() string {
	switch bm {
	case BillingMethodDirectDebit:
		return "Direct Debit"
	case BillingMethodCreditCard:
		return "Credit Card"
	case BillingMethodPayInFull:
		return "Pay in Full"
	default:
		return "Unknown"
	}
}

// BillingMethodFromFeeType converts FeeType into BillingMethod. They are equivalent
func BillingMethodFromFeeType(feeType FeeType) BillingMethod {
	switch feeType {
	case FeeTypeDirectDebit:
		return BillingMethodDirectDebit
	case FeeTypeCreditCard:
		return BillingMethodCreditCard
	case FeeTypePayInFull:
		return BillingMethodPayInFull
	}
	return BillingMethodUnknown
}

type RecurringFeeRequest struct {
	AdminFee        decimal.Decimal `json:",omitempty"`
	ContactId       int64
	CycleFee        decimal.Decimal `json:",omitempty"`
	FirstBillDate   *time.Time      `json:",omitempty"`
	ProrateFee      decimal.Decimal `json:",omitempty"`
	RecurringFeeId  int64
	RenewalDate     *time.Time `json:",omitempty"`
	SalespersonId   int64
	StartDate       time.Time
	TerminationDate *time.Time `json:",omitempty"`
}

// PostRecurringFee adds a Bolt-on to an existing membership
func (this *ApiClient) PostRecurringFee(request *RecurringFeeRequest) (target string, err error) {
	target, err = this.PostApiRequestRawResponse("/Joining/RecurringFee", nil, request)

	return
}
