package legend

import (
	"net/url"
)

type ContactWithEmailResultModel struct {
	ContactId int64
	// ExpiryDate DateTime
	FirstName        string
	IsFreeMembership bool
	IsProspect       bool
	LastName         string
	LinkedStatus     string
	MemberNumber     string
	MemberStatus     string
	Mobile           string
	Phone            string
	PostCode         string
	// ProspectCreationDate 	DateTime
	ProspectId     int64
	ProspectStatus string
}

// A prospect is assumed as active if it has not been contacted for last few months.
// 'Number of months' is a configurable setting in Legend.
func (this *ApiClient) GetContactWithEmail(email, firstName, lastName string) (target []ContactWithEmailResultModel, err error) {
	params := url.Values{}

	params.Add("Email", email)
	params.Add("FirstName", firstName)
	params.Add("LastName", lastName)
	params.Add("RetrieveMembers", "true")

	err = this.GetApiRequest("/Joining/ContactWithEmail", params, &target)

	return
}
