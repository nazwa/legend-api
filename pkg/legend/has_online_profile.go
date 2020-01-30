package legend

import (
	"net/url"
)

type HasOnlineProfile struct {
	Id           int64
	DOB          LegendTime
	IsLinked      bool
	LasName      string
	MemberNumber string
	MemberStatus string
}

// GetHasOnlineProfile checks whether the given email already has an online profile
func (this *ApiClient) GetHasOnlineProfile(email string) (target *HasOnlineProfile, err error) {
	params := url.Values{}

	params.Add("emailAddress", email)

	err = this.GetApiRequest("/Joining/HasOnlineProfile", params, &target)

	return
}
