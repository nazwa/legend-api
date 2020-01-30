package legend

import (
	"time"

	"github.com/shopspring/decimal"
)

type Gender int64
type MembershipProcessType int64

const (
	GenderMale      Gender = 1
	GenderFemale    Gender = 2
	GenderCorporate Gender = 3
	GenderUnknown   Gender = 4

	MembershipProcessTypeNew     MembershipProcessType = 0
	MembershipProcessTypeRenewal MembershipProcessType = 1
)

type MandateInfo struct {
	AccountName   string `json:",omitempty"`
	AccountNumber string `json:",omitempty"`
	BillDay       int64  `json:",omitempty"`
	Sortcode      string `json:",omitempty"`
}

type PaymentInfo struct {
	AuthorizationCode string    `json:",omitempty"`
	CardNo            string    `json:",omitempty"`
	DateProcessed     time.Time `json:",omitempty"`
	TransactionId     string    `json:",omitempty"`
}

type PersonalInfo struct {
	AddressCity                string  `json:",omitempty"`
	AddressLine1               string  `json:",omitempty"`
	AddressLine2               string  `json:",omitempty"`
	AddressLine3               string  `json:",omitempty"`
	AddressRegios              string  `json:",omitempty"`
	Birthday                   string  `json:",omitempty"` // 10-01-2010
	CommunicationPreferenceIds []int64 `json:",omitempty"`
	EmailAddress               string  `json:",omitempty"`
	EmergencyContact           string  `json:",omitempty"`
	EmergencyPhone             string  `json:",omitempty"`
	Employer                   string  `json:",omitempty"`
	EthnicOrigin               int64   `json:",omitempty"`
	FirstName                  string  `json:",omitempty"`
	GenderId                   Gender  `json:",omitempty"`
	InductionOptout            bool    `json:",omitempty"`
	MarketingSourceId          int64   `json:",omitempty"`
	MedicalConditionIds        string  `json:",omitempty"`
	MobilePhoneNumber          string  `json:",omitempty"`
	PhoneNumber                string  `json:",omitempty"`
	Postcode                   string  `json:",omitempty"`
	Surname                    string  `json:",omitempty"`
	Title                      string  `json:",omitempty"`
	WorkPhone                  string  `json:",omitempty"`
}

type PurchaseInfo struct {
	FirstBillDate         time.Time       `json:",omitempty"`
	LocationId            int64           `json:",omitempty"`
	MembershipEndDate     time.Time       `json:",omitempty"`
	MembershipStartDate   time.Time       `json:",omitempty"`
	MembershipTypePriceId int64           `json:",omitempty"`
	PromotionCode         string          `json:",omitempty"`
	PromotionId           int64           `json:",omitempty"`
	TotalFee              decimal.Decimal `json:",omitempty"`
}

type Fee struct {
	Amount              decimal.Decimal
	CardType            string `json:",omitempty"`
	MediaTypeId         int64  `json:",omitempty"`
	PaymentHandlerId    int64  `json:",omitempty"`
	ProcessType         int64
	PromotionCode       string `json:",omitempty"`
	PromotionId         int64  `json:",omitempty"`
	ServerTransactionID string `json:",omitempty"`
	SourceId            int64  `json:",omitempty"`
}

type JoiningMembersRequest struct {
	ChannelId                int64  `json:",omitempty"`
	ClientIPAddress          string `json:",omitempty"`
	ContactId                int64  `json:",omitempty"`
	DoNotCreateOnlineProfile bool   `json:",omitempty"`
	LocationID               int64  `json:",omitempty"`
	MemberNo                 string `json:",omitempty"`
	MembershipProcessType    MembershipProcessType
	PAndAEligibilities       string        `json:",omitempty"`
	ProspectId               int64         `json:",omitempty"`
	SalesPersonId            int64         `json:",omitempty"`
	SignUpLocation           string        `json:",omitempty"`
	TermsAccepted            bool          `json:",omitempty"`
	Fees                     []Fee         `json:",omitempty"`
	MandateInfo              *MandateInfo  `json:",omitempty"`
	PaymentInfo              *PaymentInfo  `json:",omitempty"`
	PersonalInfo             *PersonalInfo `json:",omitempty"`
	PurchaseInfo             *PurchaseInfo `json:",omitempty"`
}

type JoiningMembersResponse struct {
	Success             bool
	FailureReason       string
	MemberNumber        string
	MandateReference    string
	ExternalId          string
	OnlinePurchaseId    int64
	ContactId           int64
	StatusCode          int64
	ReasonPhrase        string
	RequestMessage      string
	PasswordUrl         string
	IsSuccessStatusCode bool
}

// PostJoiningMembers creates a new member
func (this *ApiClient) PostJoiningMembers(request *JoiningMembersRequest) (target JoiningMembersResponse, err error) {
	err = this.PostApiRequest("/Joining/Members", nil, request, &target)

	return
}
